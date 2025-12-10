# Uživatelský manuál

Aplikace Linkra umožňuje ukládání webových zdrojů do webového archivu a vytvoření jejich citačního záznamu, který zahrnuje archivní URL adresu. Je navržena k archivaci stránek/dokumentů, jako je například odborný článek (1 page/1 URL). Pro hlubší archivaci celého média (více URL na 1 zdroj) je možné navrhnout web k archivaci prostřednictvím formuláře do českého [Webarchivu](https://www.webarchiv.cz/cs/pridat-web).

## Postup práce s aplikací:

Aplikace se nalézá na URL adrese: http://linkra.webarchiv.cz/. Pokud chce uživatel zdroj archivovat ve webovém archivu, postupuje podle jednotlivých kroků z **Hlavní strany**. Pokud chce vytvořit citační záznam již existující archivní verze, může přejít přímo do **Generátoru citací**.

Práce spočívá ve třech základních krocích.

## 1. Vložení jedné nebo více URL adres

Uživatel vloží do formuláře jednu (nebo více) URL adres a stiskne tlačítko odeslat, což vede ke spuštění archivace. Limit počtu jednorázově vložených URL adres je 20. Po stisknutí tlačítka odeslat se dostane do druhého kroku, kde nalezne přehled archivovaných URL.

<img width="1440" height="900" alt="Screenshot 2025-12-10 at 14 08 36" src="https://github.com/user-attachments/assets/1a77dd75-66b5-48ec-8a44-c73d2fd4f31f" />



## 2. Stav archivace

Když se uživatel dostane do druhého kroku, může si uložit **Odkaz na stav archivace**, prostřednictvím kterého je možné se do Stavu archivace vrátit. Link lze případně nalézt i v historii prohlížeče. 

<img width="1428" height="732" alt="Screenshot 2025-12-10 at 14 41 24" src="https://github.com/user-attachments/assets/82637c6b-1e51-420c-a469-f1dc905fc4f3" />



Přehled zahrnuje tři sloupce:

- **URL**
Sloupec obsahuje zdroje vložené k archivaci.

- **Zkrácená archivní URL**
Zde uživatel vidí zkrácené verze archivních URL adres, tzn. adres archivních verzí zdrojů. Zkrácenou archivní URL uživatel získá obratem. Přesměrovává na plnou archivní URL, která je však dostupná teprve až po úspěšné archivaci. Vložením zdrojů do aplikace je spuštěn proces archivace, archivní verze je však možné zobrazit až po jejich zaindexování, tedy s určitým časovým zpožděním. Archivovaný webový zdroj lze v archivu dohledat podle zkrácené i plné verze.

příklad zkrácené archivní URL:
https://linkra.webarchiv/wa/KXLJIV
<br>
příklad plné archivní URL:
https://wayback.webarchiv.cz/wayback/20251125152215/https://doi.org/10.5817/ProIn2013-2-14

Pokud není archivace zdroje ještě ukončena, zkrácená archivní URL vede na hlášku **Chyba při přesměrování**. Po úspěšné archivaci vede na archivní verzi.

<img width="697" height="198" alt="Screenshot 2025-12-10 at 14 10 01" src="https://github.com/user-attachments/assets/6859736b-35da-4150-9ee2-65d462ba780c" />


- **Stav** <br>
Stav označuje, v jakém stavu se archivace nachází. Může se jednat o tyto stavy: Čeká na archivaci, Úspěšně archivováno nebo Chyba při archivaci.

Tabulku s přehledem archivovaných zdrojů si lze stáhnout ve formátu Excel nebo CSV. 

Excel, CSV

<img width="1181" height="132" alt="Screenshot 2025-12-10 at 16 58 04" src="https://github.com/user-attachments/assets/c82bc404-c329-44e7-b147-dce555930b57" />


Tabulka zahrnuje pole: URL, Zkrácený odkaz do Webarchivu, Odkaz na detail, Stav, Odkaz do Webarchivu

- **Detail**
Prostřednictvím pole Detail se uživatel dostane na Detail archivovaného zdroje. Zahrnuje URL adresu vloženou k archivaci, stav archivace, zkrácenou archivní URL adresu, plnou archivní URL adresu a přesné datum archivace.

<img width="1427" height="725" alt="Screenshot 2025-12-10 at 14 41 40" src="https://github.com/user-attachments/assets/e9e9e335-c5fd-4ac2-97cb-1c6e3065b209" />


Stisknutím pole **Pokračovat na generátor citací** se uživatel přesune ze stránky Stav archivace do třetího kroku, kde si může vytvořit citační záznam webového zdroje.

## 3. Generátor citací

Generátor citací navazuje na předchozí kroky, lze ho ale použít i samostatně, například když chce uživatel vytvořit citační záznam k již existující archivní kopii. Pokud uživatel projde výše popsané kroky (vloží URL, spustí archivaci a následně odtud pokračuje na generátor citací prostřednictvím modrého pole Pokračovat na generátor citací), aplikace automaticky předvyplní pole – URL webového zdroje, Webový archiv, URL archivní kopie, Datum archivace a Datum citace. Pokud uživatel přistoupí rovnou do generátoru citací přes odkaz v horním menu, musí všechny údaje vyplnit manuálně. 

Do pole **Citace** se postupně propisují doplňované bibliografické údaje (metadata). Lze buď využít některou z připravených šablon, nebo strukturu citace upravit či vytvořit podle vlastních požadavků. V poli **Šablona** lze vybírat ze tří možností. První je šablona, připravená v souladu s normou **ČSN ISO 690**, druhá podle citačního stylu **APA** a třetí, označena jako **Uživatelská**, umožňuje strukturu citace upravit dle vlastních požadavků. 

<img width="730" height="235" alt="Screenshot 2025-12-10 at 14 10 32" src="https://github.com/user-attachments/assets/ca9c44b5-bca0-4343-9ad4-945584f7f6f2" />


V případě, že uživatel vloží do aplikace k archivaci více než jeden zdroj, ve třetím kroku může listovat mezi jednotlivými záznamy.



Šablona ČSN ISO 690 a APA
Pokud uživatel zvolí první nebo druhou variantu šablony, tzn. **ČSN ISO 690** nebo **APA**, může do tabulky **Metadata citace** doplnit bibliografické údaje a šablona mu je automaticky převádí do struktury v souladu s normou ČSN ISO 690 nebo podle citačního stylu APA. 

<img width="736" height="509" alt="Screenshot 2025-12-10 at 14 10 43" src="https://github.com/user-attachments/assets/13418c98-7ffd-46a6-b2b3-9e321e78f12f" />


Zahrnuje základní údaje: 

- **Autor** – uživatel má k dispozici pole Příjmení a Jméno. Pro více autorů lze využít tlačítka Přidat autora / Odebrat. Pokud je autorem korporace, lze využít pole Příjmení
- **Název webového zdroje** – například název citovaného článku
- **Název zdroje / periodikum** – textové pole určené například pro údaj o periodiku, ve kterém je článek publikován
- **Místo vydání**
- **Datum vydání** – údaj lze vyplnit ručně v požadovaném tvaru nebo využít volbu z kalendáře. V případě editování v Uživatelské šabloně lze zvolit konkrétní formát z několika nabízených možností zápisu

<img width="611" height="215" alt="Screenshot 2025-12-10 at 14 11 00" src="https://github.com/user-attachments/assets/d16a844b-7715-454c-812e-575aed8a2313" />

- **URL webového zdroje** – URL adresa živého webu
- **Webový archiv** – název webového archivu, v němž se archivní kopie nachází (například Webarchiv, Internet Archive ad.)
- **URL archivní kopie** – URL adresa archivní verze ve webovém archivu, propíše se zkrácená verze URL adresy
- **Datum archivace** – údaj je automaticky vyplněný na základě archivace realizované v 1. kroku, údaj lze libovolně upravit. Pokud uživatel využije pouze generátor citací, může údaj buď vyplnit nebo použít kalendář
- **Datum citace** – do údaje se automaticky vyplní den, kdy je citace vytvářená. Lze ho libovolně upravit. Pokud uživatel využije pouze generátor citací, může údaj buď vyplnit manuálně nebo použít kalendář
- **Textové pole** – v šabloně jsou textová pole připravená pro standardní textové údaje (např. pole Online nebo Dostupné z:). V uživatelské šabloně lze toto pole využít libovolně jak pro textové, tak i číselné informace

Pokud uživateli nabízená struktura citace nevyhovuje, může v horním modrém poli zvolit třetí možnost **(Uživatelská – Upravte existující šablonu nebo vytvořte novou)**, která mu umožní sestavit si citaci podle svých potřeb. Zobrazí se mu **Šablona**, kterou lze libovolně editovat, případně může pomocí pole **Odebrat vše** všechna navrhovaná pole zrušit a začít je znovu skládat zcela podle svých individuálních požadavků. Pořadí jednotlivých částí je možné změnit přetažením myší. Po každé úpravě šablony je potřeba stisknout tlačítko dole Použít vlastní šablonu. Teprve pak se změna promítne v citaci v modrém poli nahoře.

**Uživatelská šablona**
Uživatelská šablona je určena pro úpravu navržené struktury (stávající pole lze editovat, doplňovat, odstraňovat apod.) nebo pro tvorbu šablony vlastní. Prostřednictvím pole **Odebrat vše** a potvrzením prostřednictvím tlačítka **Použít vlastní šablonu** lze všechny pole smazat. 

<img width="343" height="558" alt="11_Linkra" src="https://github.com/user-attachments/assets/a6c2e632-859f-4b1d-91ac-fd43848e7aca" />


Pro tvorbu vlastní šablony lze využít pole **Vyberte, jakou část šablony přidat do citace**. Uživatel si pomocí nabízených polí může vytvořit strukturu podle svých vlastních požadavků. 

<img width="563" height="312" alt="Screenshot 2025-12-10 at 14 11 49" src="https://github.com/user-attachments/assets/2dcb6e6b-08fe-45d0-acc4-3ab4903e856f" />



Uživatel může ovlivnit definované parametry jednotlivých údajů – forma nebo řez písma, pořadí, interpunkce, spojky a přípony. 

V případě, že se uživatel chce vrátit do předdefinované struktury, může tak učinit prostřednictvím zvolení šablony ve vedlejším poli **ČSN ISO 690** nebo **APA**.

<img width="336" height="121" alt="Screenshot 2025-12-10 at 14 11 56" src="https://github.com/user-attachments/assets/54e5c1bc-30d1-435a-96e0-d8781a2a8507" />


Pro vykopírování jednotlivé citace můžete použít tlačítko Kopírovat. Pokud pracujete na citacích více zdrojů, lze je hromadně exportovat buď do souboru HTML s formátováním nebo bez formátování do dokumentu TXT.

<img width="697" height="194" alt="Screenshot 2025-12-10 at 14 12 04" src="https://github.com/user-attachments/assets/85c04479-cc2b-438b-b4db-d2d6c4075706" />
















**Příklady citace:**




příklad ČSN ISO 690 <br>
>BOWERS, John; STANTON, Clare a ZITTRAIN, Jonathan. What the ephemerality of the Web means for your hyperlinks. Online. In: Columbia Journalism Review. New York. 2021. Dostupné z: https://www.cjr.org/analysis/linkrot-contentdrift-new-york-times.php. Archivní kopie dostupná z: Webarchiv, http://10.3.0.122:8321/archiv/2TIIRT [arch. 2025-11-27T10:44:20+01:00]. [cit. 2025-11-27].

<img width="519" height="584" alt="Screenshot 2025-12-10 at 14 12 57" src="https://github.com/user-attachments/assets/8460ddec-36c7-49a2-be17-4d0325f077b1" />













příklad APA <br>
>BOWERS, J., STANTON, C., & ZITTRAIN, J. (2021). What the ephemerality of the Web means for your hyperlinks. Columbia Journalism Review. https://www.cjr.org/analysis/linkrot-contentdrift-new-york-times.php. Archivováno: Webarchiv, http://10.3.0.122:8321/wa/2TIIRT (2025, 27. listopadu)

<img width="501" height="588" alt="Screenshot 2025-12-10 at 14 13 12" src="https://github.com/user-attachments/assets/2e9426bd-9fc6-4d9f-a945-def2d53eacd9" />

