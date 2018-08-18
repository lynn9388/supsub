# SupSub

[![GoDoc](https://godoc.org/github.com/lynn9388/supsub?status.svg)](https://godoc.org/github.com/lynn9388/supsub)

Convert normal unicode text to superscript or subscript.

Based on official unicode consortium code chart. Includes:
- [Superscripts and Subscripts](https://www.unicode.org/charts/PDF/U2070.pdf)
- [Spacing Modifier Letters](https://www.unicode.org/charts/PDF/U02B0.pdf)
- [Phonetic Extensions](https://www.unicode.org/charts/PDF/U1D00.pdf)
- [Phonetic Extensions Supplement](https://www.unicode.org/charts/PDF/U1D80.pdf)
- [C1 Controls and Latin-1 Supplement](https://www.unicode.org/charts/PDF/U0080.pdf)

> [Superscripts and Subscripts](https://en.wikipedia.org/wiki/Superscripts_and_Subscripts)

## Install

Fist, use `go get` to install the latest version of the library:

```sh
go get -u github.com/lynn9388/supsub
```

Next, include SupSub in your application:

```go
import "github.com/lynn9388/supsub"
```

## Example

```go
fmt.Printf("x%v = %v%v = %v\n", supsub.ToSup("y"), x, supsub.ToSup("2"), math.Pow(2.0, 2.0))
```

Output:
```sh
xʸ = 2² = 4
```

## Test result and convert list

Superscripts:

(⁽ )⁾ +⁺ 0⁰ 1¹ 2² 3³ 4⁴ 5⁵ 6⁶ 7⁷ 8⁸ 9⁹ =⁼ Aᴬ Bᴮ Dᴰ Eᴱ Gᴳ Hᴴ Iᴵ Jᴶ Kᴷ Lᴸ Mᴹ Nᴺ Oᴼ Pᴾ Rᴿ Tᵀ Uᵁ Wᵂ aᵃ bᵇ cᶜ dᵈ eᵉ fᶠ gᵍ hʰ iⁱ jʲ kᵏ lˡ mᵐ nⁿ oᵒ pᵖ rʳ sˢ tᵗ uᵘ vᵛ wʷ xˣ yʸ zᶻ Æᴭ ðᶞ ŋᵑ Ǝᴲ ƫᶵ Ȣᴽ ɐᵄ ɑᵅ ɒᶛ ɔᵓ ɕᶝ əᵊ ɛᵋ ɜᶟ ɟᶡ ɡᶢ ɣˠ ɥᶣ ɦʱ ɨᶤ ɩᶥ ɪᶦ ɭᶩ ɯᵚ ɰᶭ ɱᶬ ɲᶮ ɳᶯ ɴᶰ ɵᶱ ɸᶲ ɹʴ ɻʵ ʁʶ ʂᶳ ʃᶴ ʉᶶ ʊᶷ ʋᶹ ʌᶺ ʐᶼ ʑᶽ ʒᶾ ʕˤ ʝᶨ ʟᶫ βᵝ γᵞ δᵟ θᶿ φᵠ χᵡ нᵸ ᴂᵆ ᴈᵌ ᴖᵔ ᴗᵕ ᴜᶸ ᴝᵙ ᴥᵜ ᵻᶧ ᶅᶪ −⁻

Subscripts:

(₍ )₎ +₊ 0₀ 1₁ 2₂ 3₃ 4₄ 5₅ 6₆ 7₇ 8₈ 9₉ =₌ aₐ eₑ hₕ iᵢ kₖ lₗ mₘ nₙ oₒ pₚ rᵣ sₛ tₜ uᵤ vᵥ xₓ əₔ βᵦ γᵧ ρᵨ φᵩ χᵪ −₋

Convert List:

```text
( ⁽ ₍
) ⁾ ₎
+ ⁺ ₊
0 ⁰ ₀
1 ¹ ₁
2 ² ₂
3 ³ ₃
4 ⁴ ₄
5 ⁵ ₅
6 ⁶ ₆
7 ⁷ ₇
8 ⁸ ₈
9 ⁹ ₉
= ⁼ ₌
A ᴬ
B ᴮ
D ᴰ
E ᴱ
G ᴳ
H ᴴ
I ᴵ
J ᴶ
K ᴷ
L ᴸ
M ᴹ
N ᴺ
O ᴼ
P ᴾ
R ᴿ
T ᵀ
U ᵁ
W ᵂ
a ᵃ ₐ
b ᵇ
c ᶜ
d ᵈ
e ᵉ ₑ
f ᶠ
g ᵍ
h ʰ ₕ
i ⁱ ᵢ
j ʲ
k ᵏ ₖ
l ˡ ₗ
m ᵐ ₘ
n ⁿ ₙ
o ᵒ ₒ
p ᵖ ₚ
r ʳ ᵣ
s ˢ ₛ
t ᵗ ₜ
u ᵘ ᵤ
v ᵛ ᵥ
w ʷ
x ˣ ₓ
y ʸ
z ᶻ
Æ ᴭ
ð ᶞ
ŋ ᵑ
Ǝ ᴲ
ƫ ᶵ
Ȣ ᴽ
ɐ ᵄ
ɑ ᵅ
ɒ ᶛ
ɔ ᵓ
ɕ ᶝ
ə ᵊ ₔ
ɛ ᵋ
ɜ ᶟ
ɟ ᶡ
ɡ ᶢ
ɣ ˠ
ɥ ᶣ
ɦ ʱ
ɨ ᶤ
ɩ ᶥ
ɪ ᶦ
ɭ ᶩ
ɯ ᵚ
ɰ ᶭ
ɱ ᶬ
ɲ ᶮ
ɳ ᶯ
ɴ ᶰ
ɵ ᶱ
ɸ ᶲ
ɹ ʴ
ɻ ʵ
ʁ ʶ
ʂ ᶳ
ʃ ᶴ
ʉ ᶶ
ʊ ᶷ
ʋ ᶹ
ʌ ᶺ
ʐ ᶼ
ʑ ᶽ
ʒ ᶾ
ʕ ˤ
ʝ ᶨ
ʟ ᶫ
β ᵝ ᵦ
γ ᵞ ᵧ
δ ᵟ
θ ᶿ
ρ   ᵨ
φ ᵠ ᵩ
χ ᵡ ᵪ
н ᵸ
ᴂ ᵆ
ᴈ ᵌ
ᴖ ᵔ
ᴗ ᵕ
ᴜ ᶸ
ᴝ ᵙ
ᴥ ᵜ
ᵻ ᶧ
ᶅ ᶪ
− ⁻ ₋
```