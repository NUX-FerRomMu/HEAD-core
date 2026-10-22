# ADE — Pase de contexto del proyecto

## Qué es el ADE

El ADE (Agentic Development Environment) es un entorno de desarrollo orientado a coordinar trabajo de programación realizado por agentes, con el usuario como responsable de las decisiones y revisiones. Su experiencia principal es terminal-first, con interacción rápida mediante comandos y atajos compatibles con la filosofía de movimientos de Vim.

El uso habitual consiste en leer código, revisar cambios, cambiar entre tareas y atender intervenciones mientras los agentes trabajan. La unidad organizativa principal es el ticket, no el archivo abierto. La atención del usuario se organiza mediante una cola priorizada de actividades pendientes.

## Objetivo y criterio de MVP

El objetivo es alcanzar un MVP que pueda utilizarse para desarrollar el propio ADE. No se pretende una interfaz visual sofisticada al principio: importa que el flujo fundamental funcione de extremo a extremo y que la arquitectura no impida ampliar o reemplazar componentes posteriormente.

El MVP prioriza backlog, Kanban, agente de backlog, agente de código, agente de testing ejecutando pruebas y, especialmente, la priority queue de notificaciones e intervenciones.

## Decisión arquitectónica principal

Se eligió un núcleo independiente (ADE Core) con Neovim como primer frontend. Neovim aporta edición, navegación, buffers, ventanas, terminal y ecosistema de plugins; no es propietario de la lógica de negocio ni del estado de los tickets y agentes. El núcleo debe poder operar independientemente de Neovim, y otros frontends deben poder incorporarse sin rehacer el dominio.

Se prefiere una arquitectura extensible y configurable, sin exigir que toda la infraestructura de extensiones esté terminada antes del MVP. La UI gráfica podría proporcionarse mediante un frontend o plugin externo.

## Proyectos y configuración

Un proyecto se abre desde una carpeta local, como en un IDE. Cada instalación conserva su propia configuración en el entorno del usuario. Si el mismo repositorio se clona en dos computadoras, las configuraciones locales son independientes. Se busca portabilidad entre sistemas operativos; Linux es el entorno inicial previsto, no una restricción del producto.

## Backlog y Kanban

El backlog permite crear y refinar tickets. Un agente de backlog puede producir o refinar tickets y luego finalizar su ejecución. El refinamiento puede descomponer épicas o dividir tareas, conservando el conocimiento relevante para los agentes que reciban los tickets resultantes.

La asignación de tickets a agentes siempre la realiza el usuario.

Para tickets de implementación se discutió este ciclo de trabajo: `assigned → doing → committed → fixing → committed → done`. Un agente de código nace asociado a un ticket, implementa, queda disponible para correcciones o revisión mientras sea necesario y termina al cerrarse el trabajo. No existe un conjunto permanente de agentes de código esperando tickets.

Al llegar a `committed`, el flujo contempla que un agente de testing obtenga los cambios, genere o ejecute pruebas y publique sus resultados. Si falla la validación, el ticket pasa a `fixing` y el agente de desarrollo vuelve a intervenir. La aprobación de la revisión por parte del usuario equivale a autorizar el merge. Si aparecen conflictos durante la integración, puede reactivarse al agente para resolverlos.

## Agentes y proveedores

Los roles iniciales son agente de backlog, agente de código y agente de testing. Las conversaciones están asociadas a los agentes concretos de cada ticket, no a un chat genérico desvinculado del trabajo. Debe ser posible abrir el chat de un agente vivo y comunicarse con él.

El proveedor inicial preferido es Codex, utilizando su herramienta disponible, sin depender de una integración exclusiva con la API de OpenAI. Se busca independencia de proveedor para poder integrar otras herramientas o APIs en el futuro. La persistencia opcional de contexto de agentes y sus costos técnicos quedaron como asuntos de evaluación.

Las decisiones operativas deben usar heurísticas deterministas cuando sea viable; de lo contrario puede emplearse autonomía agéntica, manteniendo siempre la posibilidad de intervención humana.

## Git, edición y revisión

El ADE integra operaciones Git, ramas y revisión de cambios. La convención de ramas es configurable; se contemplan prefijos como `feat/`, `fix/` y `doc/`, por ejemplo `feat/ticket8392_unatarea`.

La revisión de código debe permitir diffs lado a lado, revisar todos o algunos archivos, comentar líneas y editar directamente. Los comentarios se transmiten al agente del ticket, que puede responder con una justificación o aplicar cambios. Si el usuario insiste, puede forzar una corrección. Las ediciones manuales del usuario producen commits propios.

Para evitar conflictos, mientras un agente modifica el trabajo no se permiten commits manuales del usuario sobre ese mismo trabajo; sí puede continuar la revisión de un commit anterior.

Se desean pruebas automáticas, coverage y controles de Git. La experiencia de edición se apoya inicialmente en Neovim.

## Priority queue

La priority queue es un componente central, visible como una columna lateral. Agrupa las intervenciones y acciones que requieren atención del usuario. Usa prioridades con aging para que los elementos postergados no queden relegados indefinidamente.

Un comando como `next` permite abrir la actividad más prioritaria. El usuario puede consultar un resumen rápido sin abandonar su contexto actual, resolver una intervención directamente cuando sea posible o navegar al contexto completo de la actividad. También puede devolver una actividad a la cola como no leída, con prioridad reducida.

El usuario conserva libertad para elegir qué atender; la cola orienta la atención, no impone una secuencia obligatoria.

## Context switching

Al cambiar a un ticket o actividad, el ADE presenta dónde quedó el trabajo: resumen de contexto, estado actual, acciones pendientes y recursos pertinentes, incluidos conversación, cambios y resultados. El resumen debe permanecer accesible. La navegación rápida entre Kanban, chats, diffs, pruebas y notificaciones es parte esencial de la experiencia.

## Control humano y configuración

El usuario puede realizar manualmente cualquier acción que el sistema automatice. Las decisiones de asignación son humanas; la aprobación de revisión es humana. Los comandos, atajos, comportamiento y mecanismos de extensión deben ser configurables, inspirados en la flexibilidad de Neovim.

## Estado actual del proyecto

La visión, los principales flujos funcionales y la decisión arquitectónica están establecidos. Se está elaborando la especificación funcional. No hay un repositorio implementado identificado en estas conversaciones. Quedan detalles de contratos e implementación por concretar; este documento describe el contexto del producto y no pretende ser una especificación técnica cerrada.
