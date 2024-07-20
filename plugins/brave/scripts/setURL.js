const chromium = Application("Brave");

const windows = chromium.windows();

for (const window of windows) {
	const tabs = window.tabs();
	for (const tab of tabs) {
		if (tab.id() === "%s") {
			tab.url = "%s";
		}
	}
}
