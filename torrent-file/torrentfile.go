package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

// Open parses a torrent file
func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return TorrentFile{}, err
	}
	defer file.Close()
	bto := bencodeTorrent{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		return TorrentFile{}, err
	}
	return bto.toTorrentFile()
}

func (i *bencodeInfo) hash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *i)
	if err != nil {
		return [20]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}

func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error) {
	infoHash, err := bto.Info.hash()
	if err != nil {
		return TorrentFile{}, err
	}

	if len(bto.Info.Pieces)%20 != 0 {
		panic("Hash not divisible by 20")
	}
	pieceHashes := make([][20]byte, 0)
	idx := 0
	for i := 0; i <= len(bto.Info.Pieces); i++ {
		var arr [20]byte

		for j := 0; j < 20; j++ {
			arr[j] = bto.Info.Pieces[idx+j]
		}

		idx += 20
		pieceHashes = append(pieceHashes, arr)
	}

	tof := TorrentFile{Announce: bto.Announce, Length: bto.Info.Length, Name: bto.Info.Name, PieceLength: bto.Info.PieceLength, PieceHashes: pieceHashes, InfoHash: infoHash}
	return tof, nil
}
