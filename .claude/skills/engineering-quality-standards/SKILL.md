---
name: engineering-quality-standards
description: Estándares de calidad de ingeniería OBLIGATORIOS para cualquier trabajo de código en este proyecto. Aplicar SIEMPRE que se escriba, modifique, revise o refactorice código (backend o frontend), y con mayor exigencia en el backend y en las consultas a base de datos. Cubre Clean Code, Clean/Hexagonal Architecture, SOLID, patrones de diseño, rendimiento de consultas, pruebas y verificación antes de afirmar.
---

## Estandares de calidad de ingenieria (OBLIGATORIO)

Toda entrega debe tener nivel de **desarrollador Java senior/experto**. No basta con que
"funcione": el codigo debe ser de calidad, legible, mantenible y eficiente. Aplicar en la
medida de lo posible (respetando el estilo existente del repo, sin reescribir de mas):
- **idioma** todo en ingles
- **Clean Code**: nombres expresivos, funciones cortas y con una sola responsabilidad,
  sin duplicacion (DRY), sin codigo muerto, sin numeros/strings magicos (usar constantes/enums).
- **Clean Architecture / Arquitectura Hexagonal**: respetar la separacion de capas
  (controller -> service -> repository, puertos/adaptadores). La logica de negocio vive en
  services, no en controllers ni repositories. No filtrar detalles de infraestructura hacia
  el dominio. Seguir la organizacion por ramo (strategy/resolver) ya existente.
- **Patrones de diseno**: usar el patron adecuado cuando aporte (Strategy para multiramo,
  Factory, Mapper, etc.) y seguir los que el proyecto ya emplea. No sobre-ingenieria.
- **SOLID** y bajo acoplamiento / alta cohesion.
- **Rendimiento (especialmente BACKEND, y con foco en consultas)**:
  - Evaluar SIEMPRE el costo de las consultas SQL/JPA. Evitar N+1, traer solo columnas
    necesarias, usar indices, paginar, y preferir una consulta correlacionada/subconsulta o
    JOIN bien pensado en vez de multiples viajes a BD. Cuidado con JOINs que duplican filas
    (usar subconsulta escalar cuando aplique).
  - Considerar transaccionalidad, timeouts (Hikari/Postgres) y posibles bloqueos.
  - Medir/razonar la complejidad antes de dar por buena una solucion.
- **Rigor**: al modificar front o back -y con mayor exigencia en el BACKEND- evaluar
  explicitamente calidad, eficiencia y rendimiento, como lo haria un senior. Mejoras fuera de
  alcance: anotarlas en el `CONTEXTO_*.md`.
- **Pruebas**: acompanar el cambio con pruebas (JUnit+Mockito / Vitest) que cubran el
  comportamiento y mantengan la cobertura. Sin comentarios en el codigo.
- **Verificar antes de afirmar**: compilar y correr las pruebas; no declarar "funciona" sin
  evidencia. Ser honesto sobre que quedo probado y que no.
