from pathlib import Path
import markdown2
from src.file_handler import FileHandler
import frontmatter

def main():

    content_path = Path('content')

    for entry in content_path.rglob("*"):
        if entry.is_dir():
            print(f"[Directory] {entry}")
            output = Path(str(entry).replace("content/", "public/"))
            output.mkdir(parents=True, exist_ok=True)
        elif entry.is_file():
            file_handler = FileHandler(entry)
            print(f"[File]      {entry}")

            static_file = file_handler.generate_static_file()

            # read template
            with open("template/blog.html", "r", encoding="utf-8") as t:
                template = t.read()

            md_content = frontmatter.load(entry)
            print(md_content.get('title'))
            post_content = md_content.content

            html_content = markdown2.markdown(post_content)
            final_content = template.replace("{{ content }}", html_content)

            with open(static_file, "w", encoding="utf-8") as f:
                f.write(final_content)

if __name__ == "__main__":
    main()