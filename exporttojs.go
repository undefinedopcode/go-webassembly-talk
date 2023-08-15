js.Global().Set(
	"emojifyMyText",
	js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		doc := js.Global().Get("document")
		textarea := doc.Call("getElementById", "my_text_area")
		if textarea.IsUndefined() {
			return nil
		}
		value := textarea.Get("value")
		if value.IsUndefined() {
			return nil
		}
		valueStr := value.String()
		emojiStr := emoji.Sprint(valueStr)
		textarea.Set("value", emojiStr)
		return nil
	}),
)

