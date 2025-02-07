// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';

let outputChannel: vscode.OutputChannel;

/**
 * Prints the given content on the output channel.
 *
 * @param content The content to be printed.
 * @param reveal Whether the output channel should be revealed.
 */
export const printChannelOutput = (content: string, reveal = false): void => {
	outputChannel.appendLine(content);
	if (reveal) {
		outputChannel.show(true);
	}
};

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
	outputChannel = vscode.window.createOutputChannel("LiteSSH");

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "litessh" is now active!');

	// The command has been defined in the package.json file
	// Now provide the implementation of the command with registerCommand
	// The commandId parameter must match the command field in package.json
	const disposable = vscode.commands.registerCommand('litessh.remoteTerminal', (resourceUri) => {
		if (!resourceUri) { return; }
		if (resourceUri.scheme === 'file') {
			const fsPath = resourceUri.fsPath;
			printChannelOutput('fsPath' + fsPath);
			const terminal = vscode.window.createTerminal(`Remote Terminal`);
			terminal.sendText(`cd "${fsPath}" && ls`);
			terminal.show();
		}

		// The code you place here will be executed every time your command is executed
		// Display a message box to the user
		vscode.window.showInformationMessage('Hello World from LiteSSH!');
	});

	context.subscriptions.push(disposable);
}

// This method is called when your extension is deactivated
export function deactivate() { }
