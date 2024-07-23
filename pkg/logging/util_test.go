package logging

import "encoding/json"

type logStrings struct {
	data map[string]string
}

func (l *logStrings) Write(p []byte) (int, error) {
	err := json.Unmarshal(p, &l.data)
	if err != nil {
		l.data = nil
		return 0, err
	}

	return len(p), nil
}

type logResult struct {
	data map[string]any
}

func (l *logResult) Write(p []byte) (int, error) {
	err := json.Unmarshal(p, &l.data)
	if err != nil {
		l.data = nil
		return 0, err
	}

	return len(p), nil
}
