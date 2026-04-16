🎬 TMDB Movie CLI (Golang)

Bu proje, The Movie Database (TMDB) API kullanarak terminal üzerinden film listelerini çekmenizi sağlayan basit bir Golang CLI uygulamasıdır.

Kullanıcı, komut satırı argümanı ile hangi tür filmleri görmek istediğini seçebilir.

Özellikler

Popüler filmleri listeleme
Vizyondaki (Now Playing) filmleri listeleme
En yüksek puanlı filmleri listeleme
Yakında çıkacak filmleri listeleme
.env üzerinden güvenli API anahtarı kullanımı
Basit ve genişletilebilir yapı

Kurulum
1. Projeyi klonla
git clone https://github.com/kullaniciadi/tmdb-cli.git
cd tmdb-cli
2. Bağımlılıkları yükle
go mod tidy
3. .env dosyası oluştur

Proje dizinine .env dosyası ekle:

TMDB_API_KEY=your_api_key_here

API anahtarını almak için:
https://www.themoviedb.org/settings/api

Kullanım

Uygulamayı çalıştırırken -type parametresi ile film türünü seçebilirsin:

go run main.go -type=popular

Desteklenen Türler

Parametre	Açıklama

popular	Popüler filmler
playing	Şu an vizyonda olanlar
top	En yüksek puanlı filmler
upcoming	Yakında çıkacak filmler

Örnek Çıktı

popular Filmler:

Baslik: Inception
Puani: 8.80
Cikis Tarihi: 2010-07-16
---------

Baslik: Interstellar
Puani: 8.60
Cikis Tarihi: 2014-11-07
---------
Proje Yapısı
.
├── main.go
├── .env
├── go.mod
└── README.md

Kullanılan Teknolojiler

Golang
net/http
encoding/json
godotenv
Geliştirme Fikirleri
Film arama özelliği ekleme
Tür (genre) filtreleme
Çoklu dil desteği
Sonuçları dosyaya kaydetme (JSON/CSV)
Poster URL'lerini gösterme
Notlar
API anahtarını kesinlikle paylaşma
.env dosyasını .gitignore içine ekle
Lisans

Bu proje eğitim amaçlıdır.
