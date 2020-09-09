# TextAnalysis

Koodi ja õigustekstide masinanalüüsi POC. Repos on järgmised katsetused:

1  Identifikaatorite sageduse ja ulatuse analüüs koodis (kaust `count`).

Vt [https://www.fluentcpp.com/2018/10/09/3-things-that-counting-words-can-reveal-on-your-code/](https://www.fluentcpp.com/2018/10/09/3-things-that-counting-words-can-reveal-on-your-code/ ).

2  Valimisseaduste ühisosa ja erinevused

[KOVVS+RKVS.html](https://e-gov.github.io/CodeCount/KOVVS+RKVS.html) - kohaliku omavalitsuse valimise seaduse (KOVVS) ja Riigikogu valimise seaduse (RKVS) tekstide võrdlus. Ainult KOVVS-is esinev tekst on välja toodud punasega, ainult RKVS-s esinev tekst rohelisega. Seadused on 01.01.2021 kehtima hakkavas redaktsioonid.

Tekstide ühitamisel on kasutatud Google vahendit [https://github.com/google/diff-match-patch](https://github.com/google/diff-match-patch).

3  Valimisseaduste sõnavara

[SÕNAD.txt](https://github.com/e-gov/CodeCount/law/blob/master/S%C3%95NAD.txt) sisaldab Eesti nelja valimisseaduse (RKVS, KOVVS, EPVS, RaHS) sõnade nimekirja.

[ÜhestatudKordumatud.txt](https://github.com/e-gov/CodeCount/law/blob/master/%C3%9ChtestatudKordumatud.txt) sisaldab [https://github.com/Filosoft/vabamorf](https://github.com/Filosoft/vabamorf) abil ühistatud (tüvekujule viidud) sõnade nimekirja.

