// @ts-nocheck
// Script that hanldes button clicks on the group view.
// It mainly copies stuff to clipboard
document.getElementById("copy-urls").addEventListener("click", copyColumn);
document.getElementById("copy-ids").addEventListener("click", copyColumn);

const groupLink = document.getElementById("group-link");
document
  .getElementById("copy-group-link")
  .addEventListener("click", copyFrom(groupLink));

// Event handler.
// Copy to clipboard column from the table containig group info.
function copyColumn(e) {
  let targetColumn;
  if (this.id === "copy-urls") {
    targetColumn = 1;
  } else if (this.id === "copy-ids") {
    targetColumn = 2;
  } else {
    console.error("copyColumn was called on wrong element!");
    return;
  }
  const dataCells = Array.from(
    document.querySelectorAll(
      `#group-info-table > tbody > tr > td:nth-of-type(${targetColumn}) > a`
    )
  );
  const baseURL = window.location.href.origin;
  const columnData = dataCells.reduce(
    (accumulator, currentValue) =>
      accumulator + "\n" + new URL(currentValue.href, baseURL).toString()
  );
  const result = navigator.clipboard.writeText(columnData);
  result.catch((reason) => console.error(reason));
  result.then(() => showCopied(this));
}

// Closure returning event hanlder.
// Copy url of single anchor.
function copyFrom(target) {
  function handler(e) {
    const data = target.href;
    const result = navigator.clipboard.writeText(data);
    result.catch((reason) => console.error(reason));
    result.then(() => showCopied(this));
  }
  return handler;
}

function showCopied(target) {
  const tmp = target.textContent;
  target.textContent = "Zkopírováno";
  setTimeout(() => (target.textContent = tmp), 1000);
}

// Reload the page after some time to fetch changes from the server
const captureCompleted = document.getElementById("capture-completed");
if (captureCompleted) {
  // If the capture is done, do nothing.
  console.log("Capture is already done. Autoreload is disabled.");
} else {
  const timeoutMs = 30 * 1000;
  window.setTimeout(() => reloadWhenVisible(), timeoutMs);
  console.log(
    "Page will reload after",
    timeoutMs / 1000,
    "seconds to fetch new data."
  );
}

// This should only reload the page if it is already visible or after it becomes visible.
function reloadWhenVisible() {
  if (!document.hidden) {
    location.reload();
  } else {
    document.addEventListener("visibilitychange", () => {
      if (!document.hidden) {
        location.reload();
      } else {
        // In some weird case that we are still hidden wait for a few seconds before retrying
        window.setTimeout(() => reloadWhenVisible(), 5 * 1000);
      }
    });
  }
}
