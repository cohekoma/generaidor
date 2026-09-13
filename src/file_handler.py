from pathlib import Path

class FileHandler:
    def __init__(self, file_entry):
        self.file_entry = file_entry

    def generate_static_file(self):
        dest_path = Path(str(self.file_entry).replace("content/", "public/").replace(".md", ".html"))
        dest_path.touch(exist_ok=True)

        return dest_path