from textnode import TextNode, TextType
def main():
    my_text = TextNode("This is some anchor text", TextType.PLAIN, "https://bc.com")
    print(my_text)

main()