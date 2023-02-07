---
## Estudo de caso do projeto
Nome do Projeto: Code Delivery
- Sistema de entregas que permite visualizar em tempo real o veículo do entregador.
- Há a possbilidade de múltiplos entregadores simultâneos.
- Serviço simulador que enviará a posição em tempo real de cada entregador.
	- Poderia ser substituído por um app mobile para o usuário final
- Os dados de cada entrega, bem como as posições, serão armazenadas no ElasticSearch para futuras análises
## Alguns desafios
- Para evitar perda de informação caso o serviço backend fique indisponível por alguns -momentos, NÃO, trabalharemos com REST.
	- Solução: Trabalharemos com o Apache Kafka para o envio e recebimento de dados entre os sistemas
- Não é responsabilidade do serviço backend persistir os dados no Elasticsearch. Logo. como armazenar as informações no Elasticsearch?
	- Solução: Utilizaremos o Kafka Connect que também consumirá os dados do simulador e fará a inserção no Elasticsearch
- Precisaremos exibir em tempo real a localização de cada entregador
	- Solução: Trabalharemos com websockets. O backend receberá os dados do simulador. e enviará as posições para o front via websocket.
## Dinâmica do sistema
print lá por 34 minutos de vídeo
## Tecnologias a serem utilizadas
- Simulador: Golang
	Por conta do multi threading
- Backend: Nest.js & Mongo
- Frontend: React
- Kafka & Kafka Connect
- Elasticsearch & Kibana
- Docker e Kubernetes
- Istio, Kiali, Prometheus & Grafana
## Bora dar uma espiada no sistema =)
print lá por 39 minutos de vídeo
## Introdução ao Apache Kafka
### Event-driven
Por exemplo:
- Carros
- E-commerce
- Alarmes
- Monitoramento
- Microsserviços
Seria muito complicado se dependessemos de uma requisição síncrona para lidar com tudo isso de uma vez só. É pra isso que o Kafka surge, ele ajuda em qualquer tipo de coisas que tenha eventos, transações, transformação.
### Tempo real
Falar em Kafka é falar em tempo real. Ele tem uma latência muito baixa.
### Histórico dos dados
O Kafka consegue manter um histórico dos dados em disco. Diferente de um RabbitMQ que executa e depois deleta os dados. O Kafka mantém e você pode setar o tempo de persistência que você quer para ele. 
### Características
- Ele não é uma sistema simples.
- O Wesley diria que ele é uma plataforma, pois ao redor dele você tem todo um ecossistema ao redor. 
- Com o Kafka você trabalha de forma distribuída.
- Tudo acaba se baseando no Kafka
	- Não pode perder mensagens e nem ficar fora do ar
- Possui um banco de dados
- Extremamente baixo e com baixa latência
- Utiliza o disco ao invés de memória para processar os dados
### Não é apenas um sistema tradicional de filas como o RabbitMQ
- Essa confusão é gerada por um overlapping. O que esses sistemas fazem o Kafka também consegue fazer, aí o público geral confunde e acha que são iguais.
- O Kafka vai muito alem disso
## Conceitos Básicos
### Topic
- Stream que atua como um banco de dados
- Todos os dados ficam armazenados, ou seja, cada Topic tem seu "local" para armezenar seus dados.
- Tópico possui diversas partições
### Partições
- É aqui que muitas pessoas bugam a mente
- O Kafka não coloca todos os ovos na mesma cesta
- Cada partição é definido por número. Ex: 0, 1, 2...
- Você é obrigado a definir a quantidade de partições quando for criar um Topic

![[Pasted image 20230206202257.png]]
		
Vamos imaginar um tópico de vendas. Esse tópico tem todas as suas partições. Conforme as mensagens vão sendo enviadas para esse tópico, cada mensagem vai caindo em uma partição diferente e você vê que eles vão criando como se fossem índices. Esses índices são chamados de offsets. Se você for ver, essas mensagens possuem seus próprios offsets.
		
Então, isso ajudaria no dia a dia no exemplo de um ecommerce. Por exemplo: quando você está no meio de uma compra e tem uma queda de rede ou um servidor fica fora do ar, coisas desse tipo. O ideal não é que você perca todo o seu carrinho de compras por conta de um erro desses, assim, as outras partições irão possuir os dados da sua compra*

A grande sacada é que ao invés de enviar a sua mensagem diretamente para o serviço que vai processar, você envie para o Kafka, e assim, aquele serviço vai ler do Kafka quando tudo estiver disponível.

Ex: Se seu sistema de cartão de crédito está fora do ar na hora da transação, não tem problema. A transação vai cair no Kafka e quando o sistema de cartão voltar, ele consome aquele tópico do Kafka e processa a mensagem.
### Producer: todo cara que manda alguma mensagem para um tópico do Kafka
### Consumer: todo cara que vai ler a mensagem do Kafka é chamado de consumer
- Pode ter vários consumers para o mesmo producer.

![[Pasted image 20230206203441.png]]

### Kafka Cluster
- O Kafka no final das contas é um cluster. 

![[Pasted image 20230206203944.png]]

- Um cluster é um conjunto de Brokers do Kafka.
### Broker
- É uma máquina onde o Kafka fica instalado. 
- Praticamente, cada broker é um servidor.
- Cada broker é responsável por armazenar os dados de uma partição. 
- Cada partição de Topic está distribuído em diferentes brokers
### Replication Factor
- Garante que pelo menos duas cópias vão estar disponíveis entre os Brokers
- Se uma máquina cair você consegue distruibuir todo esse risco
### Líderes e Followers
### Apache ZooKeeper
- É um serviço de Service Discovery que vai ficar o tempo inteiro verificando aquela máquina. Se ele ver que a máquina não está saudável, ele vai fazer o rebalanceamento de responsabilidades e define qual vai ser o novo líder.
### Consumer Group

	![[Pasted image 20230206204754.png]]
	
### Ecossistema
- Kafka Connect
	- Connectors
	Através de um connector eu consigo pegar dados de um sistema e jogar para outros. Por exemplo, colocar o meu connector para ficar assistindo um servidor MySQL e que toda vez que houver uma mudança nesse banco ele jogue essa alteração no Kafka. Aí um outro connector, chamado de sync, ele vai pegar essa mensagem do Kafka e jogar no MongoDB
- Confluent Schema Registry
- Rest Proxy
- ksqlDB
- Streams
# Colocando a mão na massa
## Go
### Go mod
O Go possui o Go Mod, que funciona como um controlador de gerenciamento de versão de todos os pacotes externos que você utilizar no Go.
