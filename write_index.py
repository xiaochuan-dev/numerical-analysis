import os

directory = 'output'
html_files = [file for file in os.listdir(directory) if file.endswith('.html')]
a_str = ''

for html_file in html_files:
    a_str += f'<li><a href="/{html_file}">{html_file}</a></li>'

html_content = f'''
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <ul>
        {a_str}
    </ul>
</body>
</html>
'''

with open('./output/index.html', 'w', encoding='utf-8') as file:
    file.write(html_content)

print("HTML 文件已成功写入。")