package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Translator struct {
	ApiKey string
}

func NewTranslator() *Translator {
	return &Translator{
		ApiKey: os.Getenv("OPENROUTER_API_KEY"),
	}
}

func (t *Translator) Translate(text string) (string, error) {
	if t.ApiKey == "" {
		// Mock refinement if no API key
		return "[MOCK REFINED] " + text, nil
	}

	url := "https://openrouter.ai/api/v1/chat/completions"
	systemPrompt := `Anda adalah seorang editor sastra profesional dan pakar lokalisasi novel. 
Tugas utama Anda: Mengubah teks Bahasa Indonesia yang terasa seperti "terjemahan mesin kaku" menjadi prosa yang "Sangat Luwes, Natural, dan Mengalir" layaknya novel terbitan profesional.

Kriteria Kualitas Utama (CRITICAL):
1. ADAPTASI IDIOMATIK: Jangan terjemahkan frasa asing secara literal. 
   - Contoh Buruk: "Apa yang tidak akan saya berikan untuk..." (Literal dari "What I wouldn't give for").
   - Contoh Bagus: "Apa pun akan saya lakukan demi...", "Rasanya aku rela memberikan apa saja untuk...", atau "Betapa inginnya aku mencicipi...".
2. HILANGKAN STRUKTUR ASING: Hindari pola kalimat pasif bahasa Inggris yang dipaksakan. Ubah menjadi pola kalimat aktif atau pasif Bahasa Indonesia yang alami.
3. KOSAKATA KAYA: Gunakan diksi yang variatif. Ganti kata umum (melihat, berjalan, sedih) dengan kata yang lebih bernuansa (menatap, melangkah, pilu).
4. DIALOG HIDUP: Dialog harus terasa seperti orang Indonesia berbicara, bukan robot.
5. TANPA ARTIFAK MESIN: Hapus kata-kata sampah seperti "yang mana", "adalah seorang", "ia adalah" jika tidak diperlukan secara struktural.
6. FLOW SASTRA: Gunakan variasi panjang kalimat. Gabungkan atau pecah kalimat agar enak dibaca dengan ritme yang pas.

Output: Hanya berikan teks hasil polesan akhir. Jangan berikan penjelasan atau komentar apa pun.`

	payload := map[string]interface{}{
		"model": "google/gemini-2.0-flash-001",
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": text},
		},
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+t.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		return choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
	}

	return "", fmt.Errorf("failed to translate: %s", string(body))
}
