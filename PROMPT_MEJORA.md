# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto arrancable. Si el adjunto es una carcasa (docs/placeholders),
el asistente debe materializar la estructura del stack del briefing, sin resolver las fases del reto.

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
API REST con Go, Gin framework y GORM

### Reto
- Tema: API REST con Go, Gin framework y GORM
- Seniority: junior-l2
- Tipo: practical
- Título: Desarrollo de una API REST para gestión de usuarios
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Definición del modelo de datos — objetivo: Definir el modelo de datos para los usuarios, incluyendo atributos y relaciones. — entregable (NO resolver): Modelo de datos definido con reglas de validación.
- Fase 2: Implementación de endpoints CRUD — objetivo: Implementar los endpoints CRUD para la gestión de usuarios. — entregable (NO resolver): Endpoints CRUD implementados y funcionando correctamente.
- Fase 3: Integración con la base de datos — objetivo: Integrar la API con una base de datos para persistir los datos de los usuarios. — entregable (NO resolver): API integrada con la base de datos y funcionando correctamente.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:
// === ARCHIVO: cmd/main.go ===
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"internal/handler"
	"internal/repository"
	"internal/service"
	"pkg/validation"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err!= nil {
		log.Fatal("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, validation.NewValidator())
	userHandler := handler.NewUserHandler(userService)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	if err := router.Run(":8080"); err!= nil {
		log.Fatal("failed to start server")
	}
}

// === ARCHIVO: internal/model/user.go ===
package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
}

// === ARCHIVO: internal/handler/user_handler.go ===
package handler

import (
	"github.com/gin-gonic/gin"
	"internal/service"
	"net/http"
	"strconv"
	"internal/model"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.CreateUser(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)
	if err := h.userService.UpdateUser(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.userService.DeleteUser(uint(id)); err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusNoContent, nil)
}

// === ARCHIVO: internal/service/user_service.go ===
package service

import (
	"errors"
	"internal/model"
	"internal/repository"
	"pkg/validation"
)

type UserService struct {
	userRepo repository.UserRepository
	validator *validation.Validator
}

func NewUserService(userRepo repository.UserRepository, validator *validation.Validator) *UserService {
	return &UserService{userRepo: userRepo, validator: validator}
}

func (s *UserService) CreateUser(user *model.User) error {
	if err := s.validator.Validate(user); err!= nil {
		return err
	}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := s.validator.Validate(user); err!= nil {
		return err
	}
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

// === ARCHIVO: internal/repository/user_repository.go ===
package repository

import (
	"errors"
	"gorm.io/gorm"
	"internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err!= nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err!= nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// === ARCHIVO: pkg/validation/validation.go ===
package validation

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(s interface{}) error {
	return v.validate.Struct(s)
}

// === ARCHIVO: config/config.yaml ===
# Configuración de la aplicación
database:
  driver: sqlite
  connectionString: test.db

server:
  port: 8080

// === ARCHIVO: docs/swagger.yaml ===
swagger: "2.0"
info:
  version: "1.0"
  title: "API de Usuarios"
  description: "API para gestionar usuarios"
paths:
  /users/:
    get:
      summary: "Obtener todos los usuarios"
      responses:
        200:
          description: "Lista de usuarios"
    post:
      summary: "Crear un usuario"
      responses:
        201:
          description: "Usuario creado"
  /users/{id}:
    get:
      summary: "Obtener usuario por ID"
      responses:
        200:
          description: "Usuario encontrado"
    put:
      summary: "Actualizar usuario"
      responses:
        200:
          description: "Usuario actualizado"
    delete:
      summary: "Eliminar usuario"
      responses:
        204:
          description: "Usuario eliminado"
definitions:
  User:
    type: object
    properties:
      id:
        type: integer
        format: int64
      name:
        type: string
      email:
        type: string
      password:
        type: string
      role:
        type: string

// === ARCHIVO: test/user_test.go ===
package test

import (
	"testing"
	"internal/model"
	"internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
		t.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	user := model.User{Name: "Test User", Email: "test@example.com", Password: "password", Role: "user"}
	if err := userRepo.CreateUser(&user); err!= nil {
		t.Fatal(err)
	}
}

```
