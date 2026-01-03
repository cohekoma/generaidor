import unittest
from textnode import TextNode, TextType

class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node1 = TextNode("This is some anchor text", TextType.PLAIN, "https://bc.com")
        node2 = TextNode("This is some anchor text", TextType.PLAIN, "https://bc.com")
        self.assertEqual(node1, node2, "testing equal")

    def test_not_eq(self):
        node1 = TextNode("This is some anchor text", TextType.PLAIN, "https://bc.com")
        node2 = TextNode("This is some anchor text", TextType.BOLD, "https://bc.com")
        self.assertNotEqual(node1, node2)

    def test_is_text_type(self):
        node1 = TextNode("This is some anchor text", TextType.PLAIN, "https://bc.com")
        self.assertIsInstance(node1.text_type, TextType)

if __name__ == "__main__":
    unittest.main()