	schema "github.com/m3db/m3db/generated/proto/schema"
	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3x/time"
	err          error
) m3db.FileSetWriter {
	w.currIdx = 0
	w.currOffset = 0
	w.err = nil
	if err := openFiles(
	); err != nil {
		closeFiles(validFiles(w.infoFd, w.indexFd, w.dataFd)...)
		return err
	}
	return nil
	if w.err != nil {
		return w.err
	}

	if err := w.writeAll(key, data); err != nil {
		w.err = err
		return err
	}
	return nil
}

func (w *writer) writeAll(key string, data [][]byte) error {
func (w *writer) close() error {
	w.infoBuffer.Reset()
	return nil
}

func (w *writer) Close() error {
	err := w.close()
	if w.err != nil {
		return w.err
	}
	if err != nil {
		w.err = err
		return err
	}
	// NB(xichen): only write out the checkpoint file if there are no errors
	// encountered between calling writer.Open() and writer.Close().