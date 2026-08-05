

Gocos2D

Introducción:
	
Gocos2D es una biblioteca de desarrollo de juegos para Google Go(lang) basada en el popular 
SDK de videojuegos Cocos2D. Para aquellos que no están familiarizados con Cocos2D, la filosofía de diseño detrás
del proyecto es que tiene la intención de simplificar el desarrollo de juegos. Una forma en que lo hace es proporcionando
objetos llamados singletons que manejan los comandos de openGL y la inicialización por ti; simplificando
efectivamente la experiencia de desarrollo de juegos de manera drástica.

Aquí tienes un ejemplo de cómo es la experiencia de Cocos2D/Gocos2D en lo que respecta a los singletons.

El singleton principal se llama director. La función del director es manejar las escenas de tu juego por
ti. Puedes pasarle al director una escena que contenga todos tus sprites y la lógica relacionada con la escena;
luego el director la coloca en la pila de escenas. El director dibujará entonces tu escena en el 
contexto de openGL mientras actualiza la lógica de tu juego por ti. Si se le ordena al director cambiar
de escena por otra, por ejemplo, cuando superas el nivel 1 y necesitas comenzar el nivel 2, el director 
saca el nivel 1 de la pila. El recolector de basura luego recoge la escena extraída y el director reinicia todo el
proceso desde cero, pero esta vez con el puntero de la escena actual apuntando a la escena del nivel 2.
