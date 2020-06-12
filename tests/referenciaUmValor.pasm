    ; Esse teste soma o valor da memória com 8 ao invés de referenciar um label
    text    ; incício do código (endereço 0)
    LDR     b
    ADD     8
    STR     a
    HLT

    data    ; inicio dos dados (endereço 128)
a:  byte    0
b:  byte    5