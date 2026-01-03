from enum import Enum

class TextType(Enum):
    PLAIN = "plain text"
    BOLD = "bold text"
    ITALIC = "italic text"
    CODE = "code text"
    ANCHOR = "anchor text"
    ALT = "alt text"

class TextNode():
    def __init__(self, text, text_type: TextType, url):
        self.text = text
        self.text_type = text_type
        self.url = url
    
    def __eq__(self, value : TextNode):
        return self.url == value.url and self.text_type.name == value.text_type.name and self.text == value.text

    def __repr__(self):
        return f"TextNode({self.text}, {self.text_type.value}, {self.url})"