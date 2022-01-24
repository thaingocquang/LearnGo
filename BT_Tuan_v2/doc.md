# Đề bài
CHƠI BÀI ĐIỂM:
- Chạy server với số người chơi (vd 3), min 2, max ???
- Chia bài cho mỗi người theo thứ tự
  + Player 1: A cơ // sleep 500ms
  + Player 2: 8 bích
  + Player 3: 10 cơ
  + Player 1: A cơ - K cơ
  + Player 2: 8 bích - 3 chuồn
  + ...
  + Player 3: 10 cơ - 4 rô - 2 bích
- Kết quả:
  + Player 1: 5 điểm
  + Player 2: 6 điểm (10 bích)
  + Player 3: 6 điểm (10 cơ) -> Thắng

# Project structure
- Card.go
    + type Card struct
    + func (c Card) IsHigher(c2 Card) bool
- DeckCard.go
    + type DeckCard struct
    + func (dc *DeckCard) init()
    + func (dc *DeckCard) Shuffle()
- Player.go
    + type Player struct
- ListPlayers.go
    + type ListPlayers struct
    + func (lp ListPlayers) DisplayCards()
    + func (lp *ListPlayers) init(nPlayers int, shuffledDeckCard DeckCard)
    + func (lpp ListPlayers) FindWinner() (winner Player, equalRank []Player)