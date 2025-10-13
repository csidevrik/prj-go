Este capitulo es introductorio y curioso, trataremos de introducir al lector al curioso e interesante mundo de go.

Porque se pretende lo del texto anterior, pues porque veo en go un mundo de oportunidades, diferentes por ejemplo se me ocurre que si parece mas complicado que java, pero al mismo tiempo te da la oportunidad de publicar en internet, claro si es que cuentas con la cuenta en github y asi creas tu repositorio personal de ideas o de una sola idea.

Por ejemplo algo muy interesante y distinto a lo que conocia de java, porque mira que java tambien ya ha avanzado pero respecto a los import de java pues me parece genial que mientras en java se nos complica esos import aunque si que si tienes un IDE no hay problema porque te ayuda en los import, a continuacion explico la diferencia:

Mientras en java para importar una clase si tienes que saber desde que package madre viene, pues en go solo lo pones entre comillas mira la diferencia

```java
import java.util.Random ;
```

en cambio en go 

```go
import "math/rand"
```

Pero claro eso es cuando tenemos que importar solo uno, pero por lo general necesitamos mas, es aqui donde me parece interesante que con go se ve mas facil esa importada o mas descomplicada para aquello, miremos la diferencia

mientras en java

```java
import java.util.Random ;
import java.time.LocalDateTime;
```

en cambio en go 

```go
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
```

Entonces a mi modo de ver me agradaria saber mas de go

Hablemos de otras particularidades de go y es que go es un lenguaje compilado, es facil de instalar ya lo he probado en windows y linux, me falta mac pero aun no tengo plata para gastarme en una mac aunque si la probé en su momento, me siento feliz que lo hicimos 

