#!/bin/bash

# ─── Config ───────────────────────────────────────────────────────────────────
# Set the token before running:
#   export OVERM_TOKEN="eyJ..."
#   ./seed_recipes.sh
#
# Or use a .env file:
#   echo 'OVERM_TOKEN=eyJ...' > .env
#   source .env && ./seed_recipes.sh

BASE_URL="${OVERM_BASE_URL:-http://localhost:8082}"

if [ -z "$OVERM_TOKEN" ]; then
  echo "Error: OVERM_TOKEN is not set."
  echo "  export OVERM_TOKEN=\"eyJ...\""
  exit 1
fi

post_recipe() {
  local name="$1"
  local body="$2"

  echo "→ Seeding: $name"
  response=$(curl -s -o /dev/null -w "%{http_code}" \
    --location "$BASE_URL/recipe-catalog/v1/recipes" \
    --header "X-Client-Type: api" \
    --header "Content-Type: application/json" \
    --header "Authorization: Bearer $OVERM_TOKEN" \
    --data "$body")

  if [ "$response" = "201" ]; then
    echo "  ✓ Created ($response)"
  else
    echo "  ✗ Failed ($response)"
  fi
}

# ─── 1. Pozole Rojo ───────────────────────────────────────────────────────────
post_recipe "Pozole Rojo" '{
  "title": "Pozole Rojo",
  "description": "Caldo tradicional de maíz cacahuazintle con carne de cerdo y chile guajillo",
  "ingredients": [
    { "name": "Maíz cacahuazintle precocido", "quantity": 500, "unit": "g" },
    { "name": "Pierna de cerdo", "quantity": 600, "unit": "g" },
    { "name": "Chile guajillo", "quantity": 5, "unit": "units" },
    { "name": "Chile ancho", "quantity": 2, "unit": "units" },
    { "name": "Ajo", "quantity": 4, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 1, "unit": "units" },
    { "name": "Orégano seco", "quantity": 1, "unit": "tsp" },
    { "name": "Sal", "quantity": 2, "unit": "tsp" }
  ],
  "steps": [
    "Cocer la pierna de cerdo en agua con media cebolla, dos dientes de ajo y sal por 45 minutos en olla exprés",
    "Desvenar los chiles y tostarlos ligeramente en comal sin quemar",
    "Remojar los chiles en agua caliente por 15 minutos",
    "Licuar los chiles con el ajo restante, un trozo de cebolla y un poco del caldo de cocción",
    "Colar la salsa y freírla en una olla grande con un poco de aceite por 5 minutos",
    "Agregar el maíz precocido y el caldo de cerdo, incorporar la carne deshebrada",
    "Cocinar a fuego medio por 20 minutos y ajustar sal",
    "Servir con tostadas, lechuga, rábano, cebolla picada y orégano"
  ],
  "servings": 6,
  "tags": ["mexicano", "cerdo", "sopa", "tradicional", "festivo"]
}'

# ─── 2. Mole Rojo con Pollo ───────────────────────────────────────────────────
post_recipe "Mole Rojo con Pollo" '{
  "title": "Mole Rojo con Pollo",
  "description": "Mole rojo oaxaqueño con chile mulato, pasilla y chocolate amargo",
  "ingredients": [
    { "name": "Pieza de pollo", "quantity": 8, "unit": "units" },
    { "name": "Chile mulato", "quantity": 4, "unit": "units" },
    { "name": "Chile pasilla", "quantity": 3, "unit": "units" },
    { "name": "Chile ancho", "quantity": 3, "unit": "units" },
    { "name": "Chocolate amargo", "quantity": 50, "unit": "g" },
    { "name": "Jitomate", "quantity": 3, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 1, "unit": "units" },
    { "name": "Ajo", "quantity": 4, "unit": "units" },
    { "name": "Cacahuate pelado", "quantity": 50, "unit": "g" },
    { "name": "Ajonjolí", "quantity": 2, "unit": "tbsp" },
    { "name": "Tortilla frita", "quantity": 1, "unit": "units" },
    { "name": "Canela en rama", "quantity": 1, "unit": "units" },
    { "name": "Clavo de olor", "quantity": 2, "unit": "units" },
    { "name": "Caldo de pollo", "quantity": 500, "unit": "ml" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Cocer el pollo en agua con sal, cebolla y ajo hasta que esté tierno",
    "Tostar y desvenar los chiles, remojar en agua caliente 20 minutos",
    "Asar el jitomate, cebolla y ajo en comal hasta que estén suaves",
    "Freír el ajonjolí, cacahuate, canela y clavo en un poco de aceite",
    "Licuar los chiles, jitomate, cebolla, ajo, cacahuate, ajonjolí, tortilla y especias con el caldo",
    "Colar la mezcla y freírla en aceite caliente moviendo constantemente por 10 minutos",
    "Agregar el chocolate amargo en trozos y disolver completamente",
    "Incorporar el resto del caldo, ajustar sal y cocinar 20 minutos más a fuego bajo",
    "Agregar las piezas de pollo y dejar reposar 5 minutos antes de servir"
  ],
  "servings": 8,
  "tags": ["mexicano", "mole", "pollo", "oaxaca", "festivo"]
}'

# ─── 3. Quesadillas de Flor de Calabaza ──────────────────────────────────────
post_recipe "Quesadillas de Flor de Calabaza" '{
  "title": "Quesadillas de Flor de Calabaza",
  "description": "Quesadillas con masa de maíz, flor de calabaza salteada y queso Oaxaca",
  "ingredients": [
    { "name": "Masa de maíz preparada", "quantity": 400, "unit": "g" },
    { "name": "Flor de calabaza", "quantity": 200, "unit": "g" },
    { "name": "Queso Oaxaca", "quantity": 200, "unit": "g" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Ajo", "quantity": 2, "unit": "units" },
    { "name": "Chile poblano", "quantity": 1, "unit": "units" },
    { "name": "Aceite vegetal", "quantity": 1, "unit": "tbsp" },
    { "name": "Sal", "quantity": 0.5, "unit": "tsp" }
  ],
  "steps": [
    "Limpiar las flores retirando el pistilo y tallo, picarlas grueso",
    "Picar la cebolla y el chile poblano en julianas finas, machacar el ajo",
    "Calentar aceite en sartén y acitronar la cebolla y ajo por 2 minutos",
    "Agregar el chile poblano y saltear 2 minutos más",
    "Incorporar la flor de calabaza, sazonar con sal y saltear 3 minutos hasta que marchite",
    "Dividir la masa en bolas de 80g y aplanar en tortillero hasta obtener discos de 15cm",
    "Colocar queso y flor de calabaza en la mitad del disco y cerrar",
    "Cocer en comal caliente a fuego medio 3 minutos por lado hasta dorar ligeramente"
  ],
  "servings": 4,
  "tags": ["mexicano", "quesadilla", "vegetariano", "masa", "cdmx"]
}'

# ─── 4. Tacos de Carnitas ─────────────────────────────────────────────────────
post_recipe "Tacos de Carnitas" '{
  "title": "Tacos de Carnitas",
  "description": "Carnitas estilo Michoacán cocinadas en su propia grasa con naranja y especias",
  "ingredients": [
    { "name": "Pierna de cerdo con hueso", "quantity": 1000, "unit": "g" },
    { "name": "Manteca de cerdo", "quantity": 200, "unit": "g" },
    { "name": "Naranja", "quantity": 1, "unit": "units" },
    { "name": "Leche entera", "quantity": 60, "unit": "ml" },
    { "name": "Ajo", "quantity": 4, "unit": "units" },
    { "name": "Hoja de laurel", "quantity": 2, "unit": "units" },
    { "name": "Orégano seco", "quantity": 1, "unit": "tsp" },
    { "name": "Sal", "quantity": 2, "unit": "tsp" },
    { "name": "Tortilla de maíz", "quantity": 16, "unit": "units" }
  ],
  "steps": [
    "Cortar la carne en trozos grandes de 8cm, sazonar con sal",
    "Derretir la manteca en olla de fondo grueso a fuego medio",
    "Agregar la carne, el jugo de naranja, leche, ajo, laurel y orégano",
    "Cocinar a fuego medio-bajo sin tapar por 1.5 horas hasta que el líquido se evapore",
    "Subir el fuego y freír la carne en su propia grasa hasta dorar y crujir",
    "Retirar la carne, deshuesar y trozar",
    "Servir en tortillas dobles con cebolla picada, cilantro y salsa de tu preferencia"
  ],
  "servings": 8,
  "tags": ["mexicano", "cerdo", "taco", "michoacan", "calles"]
}'

# ─── 5. Enchiladas Verdes ─────────────────────────────────────────────────────
post_recipe "Enchiladas Verdes" '{
  "title": "Enchiladas Verdes",
  "description": "Tortillas bañadas en salsa de tomatillo con pollo deshebrado y crema",
  "ingredients": [
    { "name": "Tortilla de maíz", "quantity": 12, "unit": "units" },
    { "name": "Pechuga de pollo cocida y deshebrada", "quantity": 300, "unit": "g" },
    { "name": "Tomatillo", "quantity": 400, "unit": "g" },
    { "name": "Chile serrano", "quantity": 3, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Ajo", "quantity": 2, "unit": "units" },
    { "name": "Cilantro", "quantity": 20, "unit": "g" },
    { "name": "Crema ácida", "quantity": 100, "unit": "ml" },
    { "name": "Queso fresco", "quantity": 100, "unit": "g" },
    { "name": "Aceite vegetal", "quantity": 3, "unit": "tbsp" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Hervir los tomatillos y chiles en agua con sal por 10 minutos hasta que cambien de color",
    "Licuar los tomatillos, chile, cebolla, ajo y cilantro con un poco del agua de cocción",
    "Freír la salsa en aceite caliente y cocinar 8 minutos a fuego medio, ajustar sal",
    "Calentar aceite en sartén y pasar las tortillas 10 segundos por lado para suavizar",
    "Rellenar cada tortilla con pollo deshebrado y enrollar",
    "Acomodar en platón y bañar generosamente con la salsa verde caliente",
    "Decorar con crema, queso fresco desmoronado y cebolla picada"
  ],
  "servings": 4,
  "tags": ["mexicano", "pollo", "enchilada", "tomatillo", "horno"]
}'

# ─── 6. Chiles en Nogada ──────────────────────────────────────────────────────
post_recipe "Chiles en Nogada" '{
  "title": "Chiles en Nogada",
  "description": "Chile poblano relleno de picadillo con frutas, bañado en nogada y granadas",
  "ingredients": [
    { "name": "Chile poblano", "quantity": 6, "unit": "units" },
    { "name": "Carne molida de res", "quantity": 300, "unit": "g" },
    { "name": "Durazno", "quantity": 2, "unit": "units" },
    { "name": "Pera", "quantity": 1, "unit": "units" },
    { "name": "Manzana", "quantity": 1, "unit": "units" },
    { "name": "Pasas", "quantity": 50, "unit": "g" },
    { "name": "Almendra pelada", "quantity": 50, "unit": "g" },
    { "name": "Piñón", "quantity": 30, "unit": "g" },
    { "name": "Jitomate", "quantity": 2, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Nuez de castilla fresca", "quantity": 150, "unit": "g" },
    { "name": "Queso de cabra", "quantity": 100, "unit": "g" },
    { "name": "Crema ácida", "quantity": 100, "unit": "ml" },
    { "name": "Granada roja", "quantity": 1, "unit": "units" },
    { "name": "Perejil fresco", "quantity": 20, "unit": "g" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Asar los chiles directamente en la llama hasta que la piel se queme por todos lados",
    "Envolver los chiles en una bolsa de plástico 15 minutos para que suden, luego pelarlos y retirar semillas",
    "Picar la fruta en cubos pequeños, reservar",
    "Sofreír la cebolla y ajo, agregar la carne y cocinar hasta que dore",
    "Incorporar jitomate, fruta picada, pasas, almendras y piñones, sazonar y cocinar 10 minutos",
    "Licuar las nueces con queso de cabra, crema y una pizca de sal hasta obtener una salsa tersa",
    "Rellenar cada chile con el picadillo y cerrar con palillo si es necesario",
    "Bañar con nogada fría, decorar con granos de granada y perejil picado"
  ],
  "servings": 6,
  "tags": ["mexicano", "festivo", "poblano", "nogada", "temporada"]
}'

# ─── 7. Tamales Rojos de Cerdo ────────────────────────────────────────────────
post_recipe "Tamales Rojos de Cerdo" '{
  "title": "Tamales Rojos de Cerdo",
  "description": "Tamales con masa de maíz y relleno de cerdo en salsa roja de chile guajillo",
  "ingredients": [
    { "name": "Masa para tamales preparada", "quantity": 1000, "unit": "g" },
    { "name": "Manteca de cerdo", "quantity": 200, "unit": "g" },
    { "name": "Espinazo de cerdo cocido y deshebrado", "quantity": 400, "unit": "g" },
    { "name": "Chile guajillo", "quantity": 6, "unit": "units" },
    { "name": "Chile ancho", "quantity": 2, "unit": "units" },
    { "name": "Ajo", "quantity": 3, "unit": "units" },
    { "name": "Comino", "quantity": 0.5, "unit": "tsp" },
    { "name": "Hoja de maíz seca", "quantity": 30, "unit": "units" },
    { "name": "Caldo de cerdo", "quantity": 250, "unit": "ml" },
    { "name": "Sal", "quantity": 1.5, "unit": "tsp" }
  ],
  "steps": [
    "Remojar las hojas de maíz en agua caliente por 30 minutos hasta suavizar",
    "Tostar y remojar los chiles en agua caliente por 15 minutos",
    "Licuar los chiles con ajo, comino y un poco de caldo, colar y freír en aceite 5 minutos",
    "Mezclar la carne deshebrada con la salsa roja y ajustar sal",
    "Batir la manteca hasta esponjar, incorporar la masa y el caldo poco a poco hasta obtener consistencia suave",
    "Extender una capa de masa de 5mm sobre cada hoja, colocar una cucharada de relleno al centro",
    "Doblar las hojas para cerrar el tamal, acomodar de pie en vaporera",
    "Cocer al vapor a fuego medio por 1 hora y 15 minutos, verificar que la masa se despegue fácilmente"
  ],
  "servings": 15,
  "tags": ["mexicano", "tamal", "cerdo", "navidad", "vapor"]
}'

# ─── 8. Sopa de Lima ──────────────────────────────────────────────────────────
post_recipe "Sopa de Lima" '{
  "title": "Sopa de Lima",
  "description": "Caldo yucateco de pollo con lima agria, tortilla frita y chile dulce",
  "ingredients": [
    { "name": "Pechuga de pollo", "quantity": 400, "unit": "g" },
    { "name": "Lima yucateca", "quantity": 3, "unit": "units" },
    { "name": "Jitomate", "quantity": 2, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 1, "unit": "units" },
    { "name": "Chile xcatic o güero", "quantity": 1, "unit": "units" },
    { "name": "Ajo", "quantity": 3, "unit": "units" },
    { "name": "Cilantro", "quantity": 15, "unit": "g" },
    { "name": "Tortilla de maíz", "quantity": 6, "unit": "units" },
    { "name": "Aceite vegetal", "quantity": 3, "unit": "tbsp" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Cocer el pollo en 1.5 litros de agua con media cebolla, ajo y sal por 25 minutos",
    "Retirar el pollo, deshebrarlo y reservar el caldo",
    "Asar en comal el jitomate, la cebolla restante y el chile hasta que se tuesten",
    "Picar el jitomate, cebolla y chile asados, sofreír en aceite por 5 minutos",
    "Incorporar el caldo colado al sofrito y llevar a hervor",
    "Agregar el jugo de dos limas y la mitad de una lima en rodajas al caldo",
    "Cortar las tortillas en tiras y freír en aceite hasta que queden crujientes",
    "Servir el caldo con pollo deshebrado, tiras de tortilla y rodajas de lima"
  ],
  "servings": 4,
  "tags": ["mexicano", "yucatan", "sopa", "pollo", "lima"]
}'

# ─── 9. Tostadas de Tinga ─────────────────────────────────────────────────────
post_recipe "Tostadas de Tinga" '{
  "title": "Tostadas de Tinga",
  "description": "Tinga poblana de pollo con chipotle sobre tostadas con frijoles y crema",
  "ingredients": [
    { "name": "Pechuga de pollo cocida y deshebrada", "quantity": 400, "unit": "g" },
    { "name": "Jitomate", "quantity": 3, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 1, "unit": "units" },
    { "name": "Chile chipotle en adobo", "quantity": 2, "unit": "units" },
    { "name": "Ajo", "quantity": 2, "unit": "units" },
    { "name": "Aceite vegetal", "quantity": 2, "unit": "tbsp" },
    { "name": "Tostadas de maíz", "quantity": 12, "unit": "units" },
    { "name": "Frijoles refritos", "quantity": 200, "unit": "g" },
    { "name": "Crema ácida", "quantity": 80, "unit": "ml" },
    { "name": "Queso fresco", "quantity": 80, "unit": "g" },
    { "name": "Aguacate", "quantity": 1, "unit": "units" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Licuar el jitomate con el chile chipotle, ajo y un poco de agua",
    "Acitronar la cebolla en julianas en aceite caliente hasta transparentar",
    "Agregar la salsa licuada y cocinar 5 minutos a fuego medio",
    "Incorporar el pollo deshebrado, ajustar sal y cocinar 8 minutos más hasta que espese",
    "Untar frijoles refritos calientes sobre cada tostada",
    "Colocar una porción de tinga encima",
    "Terminar con crema, queso fresco, rebanadas de aguacate y salsa verde al gusto"
  ],
  "servings": 4,
  "tags": ["mexicano", "tinga", "pollo", "tostada", "puebla"]
}'

# ─── 10. Caldo Tlalpeño ───────────────────────────────────────────────────────
post_recipe "Caldo Tlalpeño" '{
  "title": "Caldo Tlalpeño",
  "description": "Caldo CDMX de pollo con garbanzo, chile chipotle y epazote",
  "ingredients": [
    { "name": "Pieza de pollo", "quantity": 4, "unit": "units" },
    { "name": "Garbanzo cocido", "quantity": 200, "unit": "g" },
    { "name": "Chile chipotle seco", "quantity": 2, "unit": "units" },
    { "name": "Jitomate", "quantity": 3, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Ajo", "quantity": 3, "unit": "units" },
    { "name": "Epazote fresco", "quantity": 3, "unit": "units" },
    { "name": "Zanahoria", "quantity": 1, "unit": "units" },
    { "name": "Ejote", "quantity": 100, "unit": "g" },
    { "name": "Aguacate", "quantity": 1, "unit": "units" },
    { "name": "Sal", "quantity": 2, "unit": "tsp" }
  ],
  "steps": [
    "Cocer el pollo en 2 litros de agua con cebolla, ajo y sal por 30 minutos",
    "Asar el jitomate y los chipotles directamente en comal hasta tostar",
    "Licuar el jitomate y los chipotles con un poco del caldo, colar y freír en aceite 5 minutos",
    "Incorporar el sofrito al caldo de pollo con el pollo dentro",
    "Agregar garbanzo, zanahoria en rodajas y ejotes partidos, hervir 10 minutos",
    "Añadir las ramas de epazote y cocinar 5 minutos más",
    "Servir en cazuela con rebanadas de aguacate y chile de agua al gusto"
  ],
  "servings": 4,
  "tags": ["mexicano", "caldo", "pollo", "garbanzo", "cdmx"]
}'

# ─── 11. Huevos Rancheros ─────────────────────────────────────────────────────
post_recipe "Huevos Rancheros" '{
  "title": "Huevos Rancheros",
  "description": "Huevos estrellados sobre tortilla con salsa ranchera de jitomate y chile",
  "ingredients": [
    { "name": "Huevo", "quantity": 4, "unit": "units" },
    { "name": "Tortilla de maíz", "quantity": 4, "unit": "units" },
    { "name": "Jitomate", "quantity": 3, "unit": "units" },
    { "name": "Chile serrano", "quantity": 2, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 0.25, "unit": "units" },
    { "name": "Ajo", "quantity": 1, "unit": "units" },
    { "name": "Aceite vegetal", "quantity": 3, "unit": "tbsp" },
    { "name": "Frijoles refritos", "quantity": 150, "unit": "g" },
    { "name": "Sal", "quantity": 0.5, "unit": "tsp" }
  ],
  "steps": [
    "Asar jitomates, chile, cebolla y ajo en comal hasta que se tuesten bien",
    "Licuar los ingredientes asados con sal sin agregar agua para una salsa espesa",
    "Freír la salsa en una cucharada de aceite caliente por 5 minutos, reservar caliente",
    "Calentar las tortillas en comal y untarles frijoles refritos",
    "Freír los huevos en aceite caliente al gusto, sin romper la yema",
    "Colocar un huevo sobre cada tortilla con frijoles",
    "Bañar con salsa ranchera caliente y servir de inmediato"
  ],
  "servings": 2,
  "tags": ["mexicano", "huevo", "desayuno", "rapido", "salsa"]
}'

# ─── 12. Birria de Res ────────────────────────────────────────────────────────
post_recipe "Birria de Res" '{
  "title": "Birria de Res",
  "description": "Birria jalisciense de res marinada en chile guajillo con achiote y especias",
  "ingredients": [
    { "name": "Chambarete de res", "quantity": 1000, "unit": "g" },
    { "name": "Chile guajillo", "quantity": 8, "unit": "units" },
    { "name": "Chile ancho", "quantity": 3, "unit": "units" },
    { "name": "Chile de árbol", "quantity": 2, "unit": "units" },
    { "name": "Pasta de achiote", "quantity": 30, "unit": "g" },
    { "name": "Ajo", "quantity": 5, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 1, "unit": "units" },
    { "name": "Jitomate", "quantity": 2, "unit": "units" },
    { "name": "Vinagre blanco", "quantity": 50, "unit": "ml" },
    { "name": "Comino", "quantity": 1, "unit": "tsp" },
    { "name": "Orégano seco", "quantity": 1, "unit": "tsp" },
    { "name": "Tortilla de maíz", "quantity": 16, "unit": "units" },
    { "name": "Sal", "quantity": 2, "unit": "tsp" }
  ],
  "steps": [
    "Tostar y remojar los chiles en agua caliente por 20 minutos",
    "Licuar los chiles con jitomate, ajo, cebolla, achiote, vinagre, comino y orégano",
    "Marinar la carne en la mezcla de chiles con sal por mínimo 2 horas, de preferencia toda la noche",
    "Colocar la carne marinada en olla exprés con toda la salsa y 500ml de agua",
    "Cocinar en olla exprés por 50 minutos, la carne debe quedar muy suave",
    "Retirar la carne, desmenuzar y reservar el consomé por separado",
    "Para quesabirria: sumergir tortillas en el consomé, rellenar con carne y queso, sellar en comal con mantequilla",
    "Servir con cebolla picada, cilantro, limón y consomé para remojar"
  ],
  "servings": 6,
  "tags": ["mexicano", "res", "birria", "jalisco", "taco"]
}'

# ─── 13. Arroz con Leche ──────────────────────────────────────────────────────
post_recipe "Arroz con Leche" '{
  "title": "Arroz con Leche",
  "description": "Postre cremoso de arroz con leche entera, canela y ralladura de limón",
  "ingredients": [
    { "name": "Arroz blanco", "quantity": 200, "unit": "g" },
    { "name": "Leche entera", "quantity": 1000, "unit": "ml" },
    { "name": "Azúcar", "quantity": 100, "unit": "g" },
    { "name": "Canela en rama", "quantity": 2, "unit": "units" },
    { "name": "Limón", "quantity": 1, "unit": "units" },
    { "name": "Canela molida", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Lavar el arroz con agua fría hasta que el agua salga clara, escurrir",
    "Cocer el arroz en 400ml de agua con las ramas de canela a fuego medio hasta absorber el agua",
    "Incorporar la leche caliente, el azúcar y la ralladura de limón",
    "Cocinar a fuego bajo moviendo frecuentemente por 25 minutos hasta que espese",
    "Retirar las ramas de canela cuando alcance la consistencia cremosa deseada",
    "Servir en tazones, espolvorear canela molida por encima",
    "Puede servirse tibio o refrigerado"
  ],
  "servings": 6,
  "tags": ["mexicano", "postre", "dulce", "leche", "canela"]
}'

# ─── 14. Sopa de Tortilla ─────────────────────────────────────────────────────
post_recipe "Sopa de Tortilla" '{
  "title": "Sopa de Tortilla",
  "description": "Sopa caldosa con tortilla frita, chile pasilla, jitomate y crema",
  "ingredients": [
    { "name": "Tortilla de maíz del día anterior", "quantity": 8, "unit": "units" },
    { "name": "Jitomate", "quantity": 4, "unit": "units" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Ajo", "quantity": 3, "unit": "units" },
    { "name": "Chile pasilla", "quantity": 2, "unit": "units" },
    { "name": "Caldo de pollo", "quantity": 1000, "unit": "ml" },
    { "name": "Aceite vegetal", "quantity": 4, "unit": "tbsp" },
    { "name": "Crema ácida", "quantity": 80, "unit": "ml" },
    { "name": "Queso Oaxaca", "quantity": 100, "unit": "g" },
    { "name": "Aguacate", "quantity": 1, "unit": "units" },
    { "name": "Sal", "quantity": 1, "unit": "tsp" }
  ],
  "steps": [
    "Cortar las tortillas en tiras de 1cm y freír en aceite caliente hasta que queden doradas y crujientes",
    "Desvenar el chile pasilla y cortarlo en aros, freír brevemente en el mismo aceite, reservar",
    "Licuar el jitomate con la cebolla y el ajo asados en comal",
    "Freír la salsa de jitomate en una cucharada de aceite por 5 minutos hasta que oscurezca",
    "Agregar el caldo de pollo caliente y hervir 10 minutos, ajustar sal",
    "Servir el caldo caliente con tiras de tortilla frita al fondo",
    "Terminar con crema, queso Oaxaca, aguacate y aros de chile pasilla"
  ],
  "servings": 4,
  "tags": ["mexicano", "sopa", "tortilla", "cdmx", "caldoso"]
}'

# ─── 15. Frijoles de Olla ─────────────────────────────────────────────────────
post_recipe "Frijoles de Olla" '{
  "title": "Frijoles de Olla",
  "description": "Frijoles negros cocidos a fuego lento con epazote y manteca de cerdo",
  "ingredients": [
    { "name": "Frijol negro seco", "quantity": 500, "unit": "g" },
    { "name": "Cebolla blanca", "quantity": 0.5, "unit": "units" },
    { "name": "Ajo", "quantity": 3, "unit": "units" },
    { "name": "Epazote fresco", "quantity": 4, "unit": "units" },
    { "name": "Manteca de cerdo", "quantity": 1, "unit": "tbsp" },
    { "name": "Sal", "quantity": 2, "unit": "tsp" }
  ],
  "steps": [
    "Lavar y remojar los frijoles en agua fría la noche anterior",
    "Escurrir los frijoles y colocarlos en olla de barro o exprés con agua que los cubra 5cm",
    "Agregar la cebolla en trozo, el ajo y la manteca de cerdo",
    "Cocer en olla exprés 40 minutos o en olla normal 2 horas a fuego bajo",
    "Cuando estén suaves agregar el epazote y la sal, nunca antes de que estén cocidos",
    "Cocinar 10 minutos más para que el epazote perfume el caldo",
    "Ajustar la cantidad de agua según preferencia, el caldo debe quedar espeso"
  ],
  "servings": 6,
  "tags": ["mexicano", "frijoles", "basico", "guarnicion", "olla"]
}'

echo ""
echo "Seeding complete."