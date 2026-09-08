# Desarrollo de una API REST para gestión de usuarios

El objetivo de este reto es desarrollar una API REST que gestione la creación, lectura, actualización y eliminación de usuarios en un sistema. La API debe ser idempotente y manejar adecuadamente los errores de validación. Los usuarios tienen atributos como nombre, email, contraseña y rol. La API debe interactuar con una base de datos para persistir los datos de los usuarios.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST con Go, Gin framework y GORM |
| **Nivel** | junior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Definición del modelo de datos

**Objetivo:** Definir el modelo de datos para los usuarios, incluyendo atributos y relaciones.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identificar los atributos necesarios para un usuario (nombre, email, contraseña, rol).
- Definir las relaciones entre usuarios y otros posibles recursos (por ejemplo, roles).
- Establecer las reglas de validación para cada atributo (por ejemplo, email único, contraseña con longitud mínima).

**Entregable:** Modelo de datos definido con reglas de validación.

<details>
<summary>Pistas de conocimiento</summary>

- Considera las restricciones de negocio al definir los atributos y relaciones.
- Piensa en cómo manejarías los errores de validación en la API.

</details>

### Fase 2: Implementación de endpoints CRUD

**Objetivo:** Implementar los endpoints CRUD para la gestión de usuarios.

**Tiempo estimado:** 4 horas

**Instrucciones:**

- Crear endpoints para crear, leer, actualizar y eliminar usuarios.
- Asegurar que los endpoints sean idempotentes.
- Manejar adecuadamente los errores de validación y devolver los mensajes correctos al cliente.

**Entregable:** Endpoints CRUD implementados y funcionando correctamente.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que los endpoints deben ser idempotentes.
- Considera cómo devolver mensajes de error claros y útiles al cliente.

</details>

### Fase 3: Integración con la base de datos

**Objetivo:** Integrar la API con una base de datos para persistir los datos de los usuarios.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Configurar la conexión a la base de datos.
- Implementar la lógica para persistir y recuperar los datos de los usuarios.
- Asegurar la consistencia de los datos y manejar los errores de la base de datos.

**Entregable:** API integrada con la base de datos y funcionando correctamente.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo manejarías los errores de la base de datos en la API.
- Piensa en cómo asegurarías la consistencia de los datos.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un modelo de datos y por qué es importante en el desarrollo de una API REST?
- **paraQueSirve**: ¿Para qué sirven los endpoints CRUD en una API REST?
- **comoSeUsa**: ¿Cómo se usa GORM para interactuar con una base de datos en Go?
- **erroresComunes**: ¿Cuáles son los errores comunes al desarrollar una API REST y cómo se pueden evitar?
- **queDecisionesImplica**: ¿Qué decisiones implica el diseño de una API REST idempotente y cómo se pueden tomar?

## Criterios de Evaluacion

- Definir correctamente el modelo de datos con reglas de validación.
- Implementar endpoints CRUD idempotentes y manejar adecuadamente los errores de validación.
- Integrar la API con una base de datos y asegurar la consistencia de los datos.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
