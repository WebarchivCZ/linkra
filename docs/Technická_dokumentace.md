# Technická dokumentace

## Požadavky

- git
- docker

### Instalace

1. Naklonování repozitáře s aplikací: `git clone https://github.com/WebarchivCZ/linkra.git`
2. 

## Spuštění jako nativní proces (pro vývoj)

### Požadavky

- linux
- git
- go 1.24.0 nebo novější
- nodejs 22.20
- npm
- valkey

## Popis aplikace

Aplikace se skládá ze tří částí:

- server - poskytuje frontend, spravuje sqlite databázi, zařazuje zdroje ke sklizení, poskytuje přesměrování na archivní kopie
- worker - sklízí zadané zdroje, získává metadata ze sklizených dat, ukládá sklizená data
- queue (fronta) - komunikace mezi komponentami aplikace

### Server

Server poskytuje uživatelské rozhraní a přímá požadavky na archivaci URL adres. Protože samotné adresy mohou být archivované více než jednou, je každé adrese přidělené jedinečné ID které poté slouží k referenci dané URL adresy z webového rozhraní (např. při generování zkrácené archivní URL nebo při zobrazení detailu), ale také při výměně metadat mezi workerem a serverem.

Po zpracování a zaznamenání URL do databáze je zařazena do fronty, odkud bude odebrána workerem ke sklizení a dalšímu zpracování. Poté co worker dokončí práci obdrží server odpověď s metadaty, která použije k vytvoření adresy archivní kopie ve waybacku (aplikaci k zobrazování archivních kopií webových stránek). Tento krok vytvoření archivní adresy je možné provést i v případě, že archivní kopie není zatím ve waybacku dostupná (např. z důvodu čekání na indexaci dat). Vygenerovaná archivní adresa odpovídá formátu nejběžnějších zpřístupňovacích aplikací používaných webovými archivy (Openwayback a PyWayback).

Server také pro každou URL určenou k archivaci vytvoří zkrácený link, který od chvíle kdy dojde k úspěšnému sklizení URL bude přesměrovávat na archivní kopii dané stránky ve waybacku webového archivu.

Server dále poskytuje rozhraní pro generování citací z archivovaných URL s možností předvyplnění některých informací pro vytvoření citace.

### Worker

Worker čte z fronty požadavky na sklizení URL adres. Aktuálně je implementován jako nodejs script s použitím sklízeče Scoop. Tento trochu netradiční nástroj (oproti např. sklízeči Heritrix který je obvykle webovými archivy používán) umožňuje Scoop rychlé sklizení jedné URL, což umožňuje urychlené generování archivních adres, protože nemusíme čekat na zaindexování všech dat.

Worker je dále zodpovědný za extrakci metadat ze skliznených dat. Script otevře a zpracuje vygenerovaný WACZ soubor a odešle metadata potřebná pro generování archivní adresy zpátky do fronty odkud si je převezme server. Worker poté uloží archivní data do specifikované cesty.

