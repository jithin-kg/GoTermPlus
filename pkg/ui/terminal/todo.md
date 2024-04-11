Designing Terminal Behavior
--------------
1. Command Input at the Bottom: A single-line entry widget where users can type commands.

2. Scrollable Output Above: A multi-line text area (read-only) that displays the command output.

3. Execute Commands on Enter: When the user presses Enter, the input command is sent for execution, and the output is displayed in the output area.

4. Command History: Allow users to navigate through previously entered commands using the up and down arrow keys.

5. Clear Command: A command or button to clear the output area.

Additional Terminal Features
-------------
1. Interactive Session Support: Handling commands that require interactive input (like sudo) might be complex and could depend on the 
2. SSH library's capabilities.

3. Text Selection and Copy: Users should be able to select text from the output area and copy it to the clipboard.

4. Color Support: If the output includes ANSI color codes, rendering them can improve readability, though this may require additional handling