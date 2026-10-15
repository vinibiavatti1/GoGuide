// Adaptar
// The adapter is a structural design pattern that allows incompatible interfaces to work together.
// It acts as a bridge between two incompatible interfaces, allowing them to communicate.

package structural

// Protocols
// The protocols below define the interfaces for writers.
type ByteWriter interface {
	WriteBytes(b []byte)
}
type TextWriter interface {
	WriteText(t string)
}

// Structs
// The structs below are concrete declarations of the writers.
type DefaultByteWriter struct {
	Source []byte
}
type DefaultTextWriter struct {
	Source string
}

// Implementations
// The functions below are implementations for the writers.
func (w *DefaultByteWriter) WriteBytes(b []byte) {
	w.Source = b
}
func (w *DefaultTextWriter) WriteText(t string) {
	w.Source = t
}

// Adapter
// The adapter is a struct that implements the TextWriter interface.
// It contains a ByteWriter and converts the text to bytes before writing it.
// This allows the TextWriter to work with any ByteWriter implementation.
type DefaultTextToByteWriterAdapter struct {
	Writer ByteWriter
}

// Adapter Implementation
// The function below is an implementation for the adapter.
// It converts the text to bytes and writes it using the ByteWriter interface.
func (w *DefaultTextToByteWriterAdapter) WriteText(t string) {
	w.Writer.WriteBytes([]byte(t))
}

// Test Adapter
// Now, we can test the adapter by creating a DefaultByteWriter and a DefaultTextToByteWriterAdapter.
// We can then use the adapter to write text to the byte writer.
func TestAdapter() {
	bw := &DefaultByteWriter{}
	adapter := DefaultTextToByteWriterAdapter{Writer: bw}
	adapter.WriteText("Hello World!")
}
