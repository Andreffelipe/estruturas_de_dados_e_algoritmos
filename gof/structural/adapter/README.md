## Padrão Adapter: Explicação e Exemplo Prático

O padrão de design Adapter (Adaptador) do Gang of Four é um padrão estrutural que permite que interfaces incompatíveis trabalhem juntas. Ele funciona como um intermediário entre duas classes, convertendo a interface de uma classe em outra interface que o cliente espera.

### Cenário Real: Sistema de Processamento de Pagamentos

No exemplo, temos um cenário comum em empresas: um sistema moderno precisa integrar-se com um sistema legado de pagamentos.

#### Problema:
- Temos um sistema legado de pagamentos (`SistemaLegado`) com métodos como `efetuarTransacao()` e `consultarSaldoDisponivel()`
- Nosso novo sistema espera trabalhar com uma interface padronizada chamada `ProcessadorPagamento` com métodos como `realizarPagamento()` e `verificarFundos()`
- As interfaces são incompatíveis, mas precisamos que funcionem juntas

#### Solução com o Padrão Adapter:
1. Criar uma classe adaptadora (`AdaptadorPagamento`) que implementa a interface esperada pelo cliente
2. Internamente, o adaptador traduz as chamadas para o formato que o sistema legado entende
3. O cliente trabalha apenas com a interface que conhece, sem saber dos detalhes de implementação

Esse é um cenário muito comum em sistemas corporativos, onde é frequente a necessidade de integrar sistemas novos com sistemas legados ou APIs de terceiros que possuem interfaces diferentes do esperado.

No código de exemplo, o adaptador permite que o `ServicoCobranca` utilize o sistema legado sem precisar modificar seu código ou conhecer os detalhes de implementação do sistema antigo.