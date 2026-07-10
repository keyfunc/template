import { app, BrowserWindow } from "electron";

const createWindow = () => {
	const win = new BrowserWindow({
		width: 800,
		height: 600,
	});

	// loadFile 是基于当前的工作目录
	win.loadFile("dist/index.html");
};

app.whenReady().then(() => {
	createWindow();
});
