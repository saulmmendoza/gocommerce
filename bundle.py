import os

base_path = '/app/gocommerce-course/_base.html'
footer_path = '/app/gocommerce-course/_footer.html'
styles_path = '/app/gocommerce-course/styles.css'
js_path = '/app/gocommerce-course/main.js'
modules_dir = '/app/gocommerce-course/modules/'

with open(base_path, 'r') as f:
    base_html = f.read()

with open(footer_path, 'r') as f:
    footer_html = f.read()

with open(styles_path, 'r') as f:
    styles_content = f.read()

with open(js_path, 'r') as f:
    js_content = f.read()

# Inline CSS
style_tag = f'<style>\n{styles_content}\n</style>'
base_html = base_html.replace('<link rel="stylesheet" href="styles.css">', style_tag)

# Inline JS
script_tag = f'<script>\n{js_content}\n</script>'
base_html = base_html.replace('<script src="main.js" defer></script>', '')
footer_html = footer_html.replace('</body>', f'{script_tag}\n</body>')

# Combine modules
modules_content = []
for i in range(1, 7):
    # Depending on how files are named, we have 01-intro.html, 02-routing.html etc.
    # Let's just list and sort them
    pass

module_files = sorted([f for f in os.listdir(modules_dir) if f.endswith('.html')])
for m_file in module_files:
    with open(os.path.join(modules_dir, m_file), 'r') as f:
        modules_content.append(f.read())

final_html = base_html + '\n'.join(modules_content) + footer_html

with open('/app/course.html', 'w') as f:
    f.write(final_html)
