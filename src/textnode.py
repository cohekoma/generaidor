from enum import Enum

class TextType(Enum):
    PLAIN = "plain text"
    BOLD = "bold text"
    ITALIC = "italic text"
    CODE = "code text"
    ANCHOR = "anchor text"
    ALT = "alt text"

class TextNode():
    def __init__(self, text, text_type, url):
        self.text = text
        self.text_type = text_type
        self.url = url
    
    def __eq__(self, value : TextNode):
        print(f"comparing: {self.text} >< {value.text}")
        return self.url == value.url and self.text_type == value.text_type and self.text == value.text

    def __repr__(self):
        return f"TextNode({self.text}, {self.text_type}, {self.url})"