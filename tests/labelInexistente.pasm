    ; Esse teste tenta referenciar um label não existente
    text    ; incício do código (endereço 0)
    LDR     b
    ADD     z
    STR     a
    HLT

    data    ; inicio dos dados (endereço 128)
a:  byte    0
b:  byte    5
c:  byte    2