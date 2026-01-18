from textnode import TextNode, TextType
from pathlib import Path
import markdown
import os
import shutil

def convert_md2html(md_):
    Path("public/index.html").absolute()
    html = markdown.markdownFromFile(md_string)
    return html

def handle_public_dir(source_dir, target_dir):
    """
    Handle the public directory by creating a new public directory if the doesn't exist. If it does, remove the existing one to clean-up and regenerate all the static files.
    """
    if not os.path.exists(target_dir):
        os.makedirs(target_dir)
    else:
        shutil.rmtree(target_dir)
        os.makedirs(target_dir)

def generate_static_files(source_dir, target_dir):
    handle_public_dir(source_dir, target_dir)

def main():
    source_dir = "content"
    target_dir = "public"
    """
    scan the source directory and convert all md files to html files in the target directory
    """
    generate_static_files(source_dir, target_dir)
    return

if __name__ == "__main__":
    main()