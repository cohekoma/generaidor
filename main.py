from pathlib import Path
import markdown2

def main():

    content_path = Path('content')
    public_path = Path('public')

    for entry in content_path.rglob("*"):
        if entry.is_dir():
            print(f"[Directory] {entry}")
            output = Path(str(entry).replace("content/", "public/"))
            output.mkdir(parents=True, exist_ok=True)
        elif entry.is_file():
            print(f"[File]      {entry}")

            dest_path = Path(str(entry).replace("content/", "public/").replace(".md", ".html"))
            dest_path.touch(exist_ok=True)

            # read template
            with open("template/blog.html", "r", encoding="utf-8") as t:
                template = t.read()

            with open(entry, "r", encoding="utf-8") as f:
                text = f.read()

            html_content = markdown2.markdown(text)
            final_content = template.replace("{{ content }}", html_content)

            with open(dest_path, "w", encoding="utf-8") as f:
                f.write(final_content)

if __name__ == "__main__":
    main()