CREATE SCHEMA "cotizaciones";

CREATE TABLE "cotizaciones"."solicitud_cotizacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "numero_solicitud" varchar(20) UNIQUE,
  "consecutivo" integer,
  "vigencia" smallint NOT NULL,
  "unidad_ejecutora_id" integer NOT NULL,
  "tipo_proceso_id" integer,
  "estado_id" integer NOT NULL,
  "titulo" text,
  "objetivo" text,
  "requisitos" text,
  "observaciones" text,
  "plazo_ejecucion" varchar(250),
  "criterio_seleccion" text,
  "plan_accion_id" integer,
  "dependencia_solicitante_id" integer,
  "dependencia_destino_id" integer,
  "ordenador_gasto_id" integer,
  "responsable_tercero_id" integer NOT NULL,
  "tipo_necesidad_id" integer,
  "tipo_contrato_id" integer,
  "medio_pago_id" integer,
  "tipo_forma_pago_id" integer,
  "area_conocimiento_id" integer,
  "nucleo_basico_id" integer,
  "fecha_apertura" date,
  "fecha_cierre" date,
  "fecha_procesamiento" timestamp,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_necesidad" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "numero_necesidad" varchar(30) NOT NULL,
  "numero_disponibilidad" varchar(30),
  "valor_disponibilidad" numeric(20,7),
  "valor_cdp" numeric(20,7),
  "rubro_codigo" varchar(50),
  "rubro_nombre" varchar(250),
  "dependencia_necesidad" varchar(250),
  "estado_necesidad" varchar(50),
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_item" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "tipo_bien_servicio_id" integer NOT NULL,
  "nombre" varchar(250) NOT NULL,
  "descripcion" text NOT NULL,
  "unidad_medida_id" integer,
  "cantidad" numeric(14,2) NOT NULL,
  "cargue_masivo_id" integer,
  "codigo_unspsc" varchar(10),
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cargue_masivo" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "plantilla_id" integer NOT NULL,
  "documento_id" integer NOT NULL,
  "nombre_archivo" varchar(255) NOT NULL,
  "cargado_por_tercero_id" integer NOT NULL,
  "total_registros" integer NOT NULL,
  "aceptados" integer NOT NULL,
  "rechazados" integer NOT NULL,
  "estado_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cargue_masivo_error" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cargue_masivo_id" integer NOT NULL,
  "hoja" varchar(30) NOT NULL,
  "fila" integer NOT NULL,
  "campo" varchar(50),
  "valor_detectado" varchar(500),
  "motivo_id" integer,
  "detalle" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_pago_parcial" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "condicion" text NOT NULL,
  "porcentaje" numeric(5,4) NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_ciiu" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "actividad_economica_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."documento_solicitud" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "tipo_documento_id" integer,
  "documento_id" integer NOT NULL,
  "nombre_archivo" varchar(255) NOT NULL,
  "extension" varchar(10) NOT NULL,
  "tamano_bytes" bigint NOT NULL,
  "cargado_por_tercero_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_requisito_documental" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "tipo_documento_id" integer NOT NULL,
  "clasificacion_id" integer NOT NULL,
  "instruccion" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_proveedor" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "proveedor_id" integer NOT NULL,
  "fecha_convocatoria" timestamp NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."observacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "proveedor_convocado_id" integer NOT NULL,
  "texto" text NOT NULL,
  "estado_id" integer,
  "fecha_vista" timestamp,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."respuesta_observacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "observacion_id" integer NOT NULL,
  "autor_tercero_id" integer NOT NULL,
  "texto" text NOT NULL,
  "publica" boolean,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cotizacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "proveedor_convocado_id" integer UNIQUE NOT NULL,
  "estado_id" integer NOT NULL,
  "fecha_primer_envio" timestamp,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cotizacion_version" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cotizacion_id" integer NOT NULL,
  "numero_version" smallint NOT NULL,
  "motivo_id" integer NOT NULL,
  "solicitud_modificacion_id" integer,
  "periodo_subsanacion_id" integer,
  "estado_version_id" integer NOT NULL,
  "fecha_vencimiento" date,
  "descuentos" text,
  "observaciones" text,
  "fecha_envio" timestamp,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cotizacion_version_item" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cotizacion_version_id" integer NOT NULL,
  "solicitud_item_id" integer NOT NULL,
  "cantidad" numeric(14,2) NOT NULL,
  "valor_unitario" numeric(20,7) NOT NULL,
  "tarifa_iva_id" integer NOT NULL,
  "porcentaje_iva" numeric(5,4) NOT NULL,
  "ficha_tecnica" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."documento_cotizacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cotizacion_id" integer NOT NULL,
  "requisito_documental_id" integer,
  "cotizacion_version_id" integer NOT NULL,
  "documento_id" integer NOT NULL,
  "nombre_archivo" varchar(255) NOT NULL,
  "extension" varchar(10) NOT NULL,
  "tamano_bytes" bigint NOT NULL,
  "reemplaza_a_id" integer,
  "cargado_por_tercero_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."periodo_subsanacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "proveedor_convocado_id" integer,
  "fecha_inicio" timestamp NOT NULL,
  "fecha_fin" timestamp NOT NULL,
  "motivo" text NOT NULL,
  "habilitado_por_tercero_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."decision_solicitante" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cotizacion_id" integer UNIQUE NOT NULL,
  "cotizacion_version_id" integer NOT NULL,
  "resultado_id" integer NOT NULL,
  "mensaje_proveedor" text NOT NULL,
  "justificacion_seleccion" text,
  "decidido_por_tercero_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."validacion_ordenador" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "decision_cotizacion_id" integer NOT NULL,
  "decision_id" integer NOT NULL,
  "observaciones" text,
  "ordenador_tercero_id" integer NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_modificacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "numero" smallint NOT NULL,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_modificacion_estado" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_modificacion_id" integer NOT NULL,
  "estado_id" integer NOT NULL,
  "responsable_tercero_id" integer,
  "justificacion" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_modificacion_detalle" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_modificacion_id" integer NOT NULL,
  "tipo_cambio_id" integer NOT NULL,
  "solicitud_item_id" integer,
  "valor_anterior" jsonb,
  "valor_nuevo" jsonb,
  "justificacion" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."solicitud_estado" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "estado_id" integer NOT NULL,
  "responsable_tercero_id" integer,
  "justificacion" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."cotizacion_estado" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "cotizacion_id" integer NOT NULL,
  "estado_id" integer NOT NULL,
  "responsable_tercero_id" integer,
  "justificacion" text,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "cotizaciones"."notificacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "solicitud_id" integer NOT NULL,
  "evento_id" integer NOT NULL,
  "proveedor_convocado_id" integer,
  "destinatario_tercero_id" integer NOT NULL,
  "canal_id" integer NOT NULL,
  "estado_envio_id" integer NOT NULL,
  "intentos" smallint NOT NULL DEFAULT 0,
  "ultimo_error" text,
  "fecha_envio" timestamp,
  "fecha_modificacion" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "cotizaciones"."solicitud_necesidad" ("solicitud_id", "numero_necesidad");

CREATE UNIQUE INDEX ON "cotizaciones"."solicitud_ciiu" ("solicitud_id", "actividad_economica_id");

CREATE UNIQUE INDEX ON "cotizaciones"."solicitud_requisito_documental" ("solicitud_id", "tipo_documento_id");

CREATE UNIQUE INDEX ON "cotizaciones"."solicitud_proveedor" ("solicitud_id", "proveedor_id");

CREATE UNIQUE INDEX ON "cotizaciones"."cotizacion_version" ("cotizacion_id", "numero_version");

CREATE UNIQUE INDEX ON "cotizaciones"."cotizacion_version_item" ("cotizacion_version_id", "solicitud_item_id");

CREATE UNIQUE INDEX ON "cotizaciones"."solicitud_modificacion" ("solicitud_id", "numero");

COMMENT ON TABLE "cotizaciones"."solicitud_cotizacion" IS 'Solicitud de cotizacion. Entidad central del modulo, una fila por solicitud, compartida por el solicitante y los proveedores convocados';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."numero_solicitud" IS 'Numero de radicado de la solicitud. Se asigna al procesarla y es unico';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."consecutivo" IS 'Consecutivo de la solicitud dentro de la vigencia';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."vigencia" IS 'Anio fiscal de la necesidad asociada a la solicitud';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."unidad_ejecutora_id" IS 'Llave foranea logica a la unidad ejecutora administrada en Kronos. Se valida por API';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."titulo" IS 'Titulo de la cotizacion, copiado de la necesidad registrada en SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."objetivo" IS 'Objetivo y justificacion de la solicitud, copiados de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."requisitos" IS 'Requisitos de la solicitud, copiados de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."observaciones" IS 'Observaciones adicionales de la solicitud, copiadas de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."plazo_ejecucion" IS 'Plazo de ejecucion del contrato, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."criterio_seleccion" IS 'Criterio con el que se seleccionara la cotizacion ganadora';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."plan_accion_id" IS 'Llave foranea logica al catalogo de planes de accion que estara en Parametros. Pendiente de validacion frente a los requerimientos';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."dependencia_solicitante_id" IS 'Llave foranea logica a la dependencia que solicita, administrada en Oikos. Se valida por API';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."dependencia_destino_id" IS 'Llave foranea logica a la dependencia destino de los bienes o servicios, administrada en Oikos. Se valida por API';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."ordenador_gasto_id" IS 'Llave foranea logica al ordenador del gasto de la solicitud. Origen por confirmar, persona en terceros';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."responsable_tercero_id" IS 'Llave foranea logica a la persona que crea la solicitud, registrada en terceros. Se valida por API';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."tipo_necesidad_id" IS 'Llave foranea logica al catalogo de tipos de necesidad que estara en Parametros. Se calcula a partir de los items';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."tipo_contrato_id" IS 'Llave foranea logica al catalogo de tipos de contrato que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."medio_pago_id" IS 'Llave foranea logica al catalogo de medios de pago que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."area_conocimiento_id" IS 'Llave foranea logica al catalogo de areas de conocimiento SNIES que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."nucleo_basico_id" IS 'Llave foranea logica al catalogo de nucleos basicos de conocimiento SNIES que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."fecha_apertura" IS 'Fecha desde la cual los proveedores convocados pueden responder la solicitud';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."fecha_cierre" IS 'Fecha de cierre de la recepcion de cotizaciones. Al menos tres dias habiles despues de la apertura';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."fecha_procesamiento" IS 'Fecha y hora en que la solicitud se proceso y se envio a los proveedores';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_necesidad" IS 'Necesidades y datos presupuestales de SICAPITAL asociados a la solicitud. Una fila por necesidad';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece la necesidad';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."numero_necesidad" IS 'Numero de la necesidad en SICAPITAL. Se identifica junto con la unidad ejecutora y la vigencia de la solicitud';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."numero_disponibilidad" IS 'Numero del certificado de disponibilidad presupuestal, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."valor_disponibilidad" IS 'Valor en pesos de la disponibilidad presupuestal de la necesidad, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."valor_cdp" IS 'Valor en pesos del certificado de disponibilidad presupuestal, copiado de SICAPITAL. Se conserva junto a valor_disponibilidad hasta confirmar si son el mismo dato';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."rubro_codigo" IS 'Codigo del rubro presupuestal de la necesidad, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."rubro_nombre" IS 'Nombre del rubro presupuestal de la necesidad, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."dependencia_necesidad" IS 'Nombre de la dependencia de la necesidad, copiado de SICAPITAL';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."estado_necesidad" IS 'Estado de la necesidad en SICAPITAL al momento de la copia';

COMMENT ON COLUMN "cotizaciones"."solicitud_necesidad"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_item" IS 'Bienes y servicios solicitados en la solicitud de cotizacion. Una fila por item';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece el item';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."tipo_bien_servicio_id" IS 'Llave foranea logica al catalogo de tipos de necesidad que estara en Parametros. Indica si el item es bien o servicio';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."nombre" IS 'Nombre del bien o servicio solicitado';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."descripcion" IS 'Descripcion y ficha tecnica del bien o servicio solicitado';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."unidad_medida_id" IS 'Llave foranea logica al catalogo de unidades de medida que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."cantidad" IS 'Cantidad solicitada del bien o servicio';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."cargue_masivo_id" IS 'Llave foranea a cotizaciones.cargue_masivo. Cargue del que proviene el item, vacio si se registro individualmente';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."codigo_unspsc" IS 'Codigo UNSPSC del bien o servicio. Por confirmar si se adopta';

COMMENT ON COLUMN "cotizaciones"."solicitud_item"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cargue_masivo" IS 'Cargues masivos de items realizados con plantilla en una solicitud. Una fila por archivo cargado';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud en la que se cargaron los items';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."documento_id" IS 'Llave foranea logica al archivo cargado en el gestor documental';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."nombre_archivo" IS 'Nombre original del archivo cargado';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."cargado_por_tercero_id" IS 'Llave foranea logica a la persona que realizo el cargue, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."total_registros" IS 'Cantidad total de registros procesados en el archivo';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."aceptados" IS 'Cantidad de registros aceptados e incorporados como items';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."rechazados" IS 'Cantidad de registros rechazados por inconsistencias';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."estado_id" IS 'Llave foranea logica al catalogo de estados del cargue masivo que estara en Parametros. Resultado del procesamiento del archivo';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cargue_masivo_error" IS 'Errores detectados al validar un cargue masivo. Una fila por error';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."cargue_masivo_id" IS 'Llave foranea a cotizaciones.cargue_masivo. Cargue en el que se detecto el error';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."hoja" IS 'Hoja de la plantilla donde esta el registro rechazado, BIEN o SERVICIO';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."fila" IS 'Numero de fila del registro rechazado en la plantilla';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."campo" IS 'Campo de la plantilla que presenta el error';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."valor_detectado" IS 'Valor encontrado en el campo con error';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."motivo_id" IS 'Llave foranea logica al catalogo de motivos de rechazo del cargue que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."detalle" IS 'Mensaje del validador que explica el error';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo_error"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_pago_parcial" IS 'Condiciones de pago de la solicitud. Una fila por condicion';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece la condicion de pago';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."condicion" IS 'Hito o entregable que habilita el pago';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."porcentaje" IS 'Proporcion del valor total que se paga en esta condicion, expresada como fraccion, por ejemplo 0.3000 para 30 por ciento. La suma por solicitud debe ser 1';

COMMENT ON COLUMN "cotizaciones"."solicitud_pago_parcial"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_ciiu" IS 'Actividades economicas CIIU elegidas para la solicitud. Una fila por actividad';

COMMENT ON COLUMN "cotizaciones"."solicitud_ciiu"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_ciiu"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_ciiu"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece la actividad economica';

COMMENT ON COLUMN "cotizaciones"."solicitud_ciiu"."actividad_economica_id" IS 'Llave foranea logica al catalogo actividad_economica que estara en Parametros. Codigo CIIU elegido';

COMMENT ON COLUMN "cotizaciones"."solicitud_ciiu"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."documento_solicitud" IS 'Anexos que adjunta el jefe de dependencia a la solicitud. Una fila por documento';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece el anexo';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."tipo_documento_id" IS 'Llave foranea logica al catalogo tipo_documento que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."documento_id" IS 'Llave foranea logica al archivo almacenado en el gestor documental';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."nombre_archivo" IS 'Nombre original del archivo anexado';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."extension" IS 'Extension del archivo anexado';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."tamano_bytes" IS 'Tamanio del archivo en bytes';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."cargado_por_tercero_id" IS 'Llave foranea logica a la persona que cargo el archivo, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."documento_solicitud"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_requisito_documental" IS 'Documentos que el oferente debe aportar para la solicitud. Una fila por documento exigido';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud que exige el documento';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."tipo_documento_id" IS 'Llave foranea logica al catalogo tipo_documento que estara en Parametros. Documento exigido al oferente';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."clasificacion_id" IS 'Llave foranea logica al catalogo de clasificacion del requisito documental que estara en Parametros. Obligatorio o complementario';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."instruccion" IS 'Instruccion para el oferente sobre el documento exigido';

COMMENT ON COLUMN "cotizaciones"."solicitud_requisito_documental"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_proveedor" IS 'Proveedores convocados a la solicitud. Una fila por proveedor';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que se convoca al proveedor';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."proveedor_id" IS 'Llave foranea logica al proveedor del registro de proveedores, con identidad en terceros. Se valida por API';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."fecha_convocatoria" IS 'Fecha y hora en que el proveedor fue convocado al procesar la solicitud';

COMMENT ON COLUMN "cotizaciones"."solicitud_proveedor"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."observacion" IS 'Observaciones de los proveedores convocados sobre la solicitud. Una fila por observacion';

COMMENT ON COLUMN "cotizaciones"."observacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."observacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."observacion"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud sobre la que se hace la observacion';

COMMENT ON COLUMN "cotizaciones"."observacion"."proveedor_convocado_id" IS 'Llave foranea a cotizaciones.solicitud_proveedor. Proveedor convocado que escribe la observacion';

COMMENT ON COLUMN "cotizaciones"."observacion"."texto" IS 'Texto de la observacion del proveedor';

COMMENT ON COLUMN "cotizaciones"."observacion"."fecha_vista" IS 'Fecha y hora en que el solicitante reviso la observacion. Vacio si no se ha revisado';

COMMENT ON COLUMN "cotizaciones"."observacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."respuesta_observacion" IS 'Respuestas de la Universidad a las observaciones de los proveedores. Una fila por respuesta';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."observacion_id" IS 'Llave foranea a cotizaciones.observacion. Observacion que se responde';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."autor_tercero_id" IS 'Llave foranea logica a la persona que responde, solicitante u ordenador, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."texto" IS 'Texto de la respuesta de la Universidad';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."publica" IS 'Indica si la respuesta es visible para todos los oferentes o solo para el autor de la observacion';

COMMENT ON COLUMN "cotizaciones"."respuesta_observacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cotizacion" IS 'Cotizacion presentada por un proveedor convocado. El contenido se guarda por versiones';

COMMENT ON COLUMN "cotizaciones"."cotizacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cotizacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cotizacion"."proveedor_convocado_id" IS 'Llave foranea unica a cotizaciones.solicitud_proveedor. Proveedor convocado que presenta la cotizacion, una por proveedor';

COMMENT ON COLUMN "cotizaciones"."cotizacion"."fecha_primer_envio" IS 'Fecha y hora en que el proveedor envio la primera version';

COMMENT ON COLUMN "cotizaciones"."cotizacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cotizacion_version" IS 'Versiones de la cotizacion del proveedor. Una version enviada no se modifica. Una fila por version';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."cotizacion_id" IS 'Llave foranea a cotizaciones.cotizacion. Cotizacion a la que pertenece la version';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."numero_version" IS 'Numero de la version dentro de la cotizacion, 1 para la inicial';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."solicitud_modificacion_id" IS 'Llave foranea a cotizaciones.solicitud_modificacion. Modificacion que origino la version, obligatoria si el motivo es adenda';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."periodo_subsanacion_id" IS 'Llave foranea a cotizaciones.periodo_subsanacion. Periodo en que se presento la version, obligatorio si el motivo es subsanacion';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."fecha_vencimiento" IS 'Fecha hasta la cual el proveedor mantiene vigente la cotizacion';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."descuentos" IS 'Descuentos ofrecidos por el proveedor';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."observaciones" IS 'Observaciones adicionales del proveedor sobre la cotizacion';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."fecha_envio" IS 'Fecha y hora en que el proveedor envio la version';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cotizacion_version_item" IS 'Precio, IVA y ficha tecnica ofrecidos por item en cada version de la cotizacion. Una fila por item y version';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."cotizacion_version_id" IS 'Llave foranea a cotizaciones.cotizacion_version. Version a la que pertenece el precio';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."solicitud_item_id" IS 'Llave foranea a cotizaciones.solicitud_item. Item de la solicitud que se cotiza';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."cantidad" IS 'Cantidad del item al momento de crear la version, copiada de la solicitud';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."valor_unitario" IS 'Valor unitario ofrecido por el proveedor antes de IVA';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."tarifa_iva_id" IS 'Llave foranea logica al catalogo de tarifas de IVA que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."porcentaje_iva" IS 'Tarifa de IVA vigente al enviar la version, expresada como fraccion, por ejemplo 0.1900 para 19 por ciento. Copiada de la tarifa';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."ficha_tecnica" IS 'Descripcion del bien o servicio ofrecido por el proveedor';

COMMENT ON COLUMN "cotizaciones"."cotizacion_version_item"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."documento_cotizacion" IS 'Documentos que adjunta el proveedor a su cotizacion. Una fila por documento';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."cotizacion_id" IS 'Llave foranea a cotizaciones.cotizacion. Cotizacion a la que pertenece el documento';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."requisito_documental_id" IS 'Llave foranea a cotizaciones.solicitud_requisito_documental. Requisito que cumple el documento, vacio si es un adjunto libre';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."cotizacion_version_id" IS 'Llave foranea a cotizaciones.cotizacion_version. Version en la que se cargo el documento';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."documento_id" IS 'Llave foranea logica al archivo almacenado en el gestor documental';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."nombre_archivo" IS 'Nombre original del archivo cargado';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."extension" IS 'Extension del archivo cargado';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."tamano_bytes" IS 'Tamanio del archivo en bytes';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."reemplaza_a_id" IS 'Llave foranea a cotizaciones.documento_cotizacion. Documento anterior que este reemplaza en una subsanacion';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."cargado_por_tercero_id" IS 'Llave foranea logica a la persona que cargo el archivo, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."documento_cotizacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."periodo_subsanacion" IS 'Periodos en los que se habilita la subsanacion de documentos de las cotizaciones. Una fila por periodo';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud en la que se habilita la subsanacion';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."proveedor_convocado_id" IS 'Llave foranea a cotizaciones.solicitud_proveedor. Proveedor al que aplica el periodo, vacio si aplica a todos';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."fecha_inicio" IS 'Fecha y hora de inicio del periodo de subsanacion';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."fecha_fin" IS 'Fecha y hora de fin del periodo de subsanacion. Debe ser posterior al inicio';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."motivo" IS 'Motivo por el que se habilita la subsanacion';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."habilitado_por_tercero_id" IS 'Llave foranea logica a la persona que habilito el periodo, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."periodo_subsanacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."decision_solicitante" IS 'Decision del solicitante sobre la cotizacion de un proveedor. La fecha de creacion es la fecha de la decision';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."cotizacion_id" IS 'Llave foranea unica a cotizaciones.cotizacion. Cotizacion sobre la que se decide, una decision por cotizacion';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."cotizacion_version_id" IS 'Llave foranea a cotizaciones.cotizacion_version. Version evaluada, debe ser la vigente';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."resultado_id" IS 'Llave foranea logica al catalogo de resultados de la decision que estara en Parametros. Aprobado o rechazado';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."mensaje_proveedor" IS 'Respuesta o recomendacion que recibe el proveedor';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."justificacion_seleccion" IS 'Razones de la seleccion del proveedor. Obligatoria si el resultado es aprobado';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."decidido_por_tercero_id" IS 'Llave foranea logica a la persona que registra la decision, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."validacion_ordenador" IS 'Validacion del ordenador del gasto sobre la decision del solicitante. Una fila por validacion';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud que se valida';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."decision_cotizacion_id" IS 'Llave foranea a cotizaciones.decision_solicitante. Decision del solicitante que se valida';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."decision_id" IS 'Llave foranea logica al catalogo de resultados de la decision que estara en Parametros. Aprueba o rechaza';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."observaciones" IS 'Observaciones del ordenador del gasto sobre la validacion';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."ordenador_tercero_id" IS 'Llave foranea logica al ordenador del gasto que valida, registrado en terceros';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_modificacion" IS 'Modificaciones solicitadas sobre una solicitud ya procesada. El estado, la justificacion y las fechas se obtienen del historial';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud que se modifica';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion"."numero" IS 'Numero de la modificacion dentro de la solicitud, el que ven los usuarios';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_modificacion_estado" IS 'Historial de estados de cada solicitud de modificacion. Una fila por cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."solicitud_modificacion_id" IS 'Llave foranea a cotizaciones.solicitud_modificacion. Modificacion a la que pertenece el cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."responsable_tercero_id" IS 'Llave foranea logica a la persona que realiza el cambio, registrada en terceros. Vacio si lo realiza el sistema';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."justificacion" IS 'Justificacion del cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_estado"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_modificacion_detalle" IS 'Detalle de los cambios de cada modificacion de la solicitud. Una fila por cambio';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."solicitud_modificacion_id" IS 'Llave foranea a cotizaciones.solicitud_modificacion. Modificacion a la que pertenece el cambio';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."solicitud_item_id" IS 'Llave foranea a cotizaciones.solicitud_item. Item modificado, obligatorio en cambios de items';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."valor_anterior" IS 'JSON con los campos del item o la fecha antes del cambio, por ejemplo nombre, descripcion, unidad, cantidad o fecha de cierre';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."valor_nuevo" IS 'JSON con los mismos campos despues del cambio';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."justificacion" IS 'Justificacion del cambio';

COMMENT ON COLUMN "cotizaciones"."solicitud_modificacion_detalle"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."solicitud_estado" IS 'Historial de estados de la solicitud. Una fila por cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."solicitud_id" IS 'Llave foranea a cotizaciones.solicitud_cotizacion. Solicitud a la que pertenece el cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."responsable_tercero_id" IS 'Llave foranea logica a la persona que realiza el cambio, registrada en terceros. Vacio si lo realiza el sistema';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."justificacion" IS 'Justificacion del cambio de estado';

COMMENT ON COLUMN "cotizaciones"."solicitud_estado"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."cotizacion_estado" IS 'Historial de estados de la cotizacion del proveedor. Una fila por cambio de estado';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."cotizacion_id" IS 'Llave foranea a cotizaciones.cotizacion. Cotizacion a la que pertenece el cambio de estado';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."responsable_tercero_id" IS 'Llave foranea logica a la persona que realiza el cambio, registrada en terceros. Vacio si lo realiza el sistema';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."justificacion" IS 'Justificacion del cambio de estado';

COMMENT ON COLUMN "cotizaciones"."cotizacion_estado"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

COMMENT ON TABLE "cotizaciones"."notificacion" IS 'Notificaciones generadas por eventos del proceso de la solicitud. Una fila por notificacion y destinatario';

COMMENT ON COLUMN "cotizaciones"."notificacion"."id" IS 'Identificador unico del registro. Llave primaria autonumerica generada por la secuencia propia de la tabla, nombre_tabla_id_seq';

COMMENT ON COLUMN "cotizaciones"."notificacion"."fecha_creacion" IS 'Fecha y hora de creacion del registro. Se asigna automaticamente al insertar';

COMMENT ON COLUMN "cotizaciones"."notificacion"."destinatario_tercero_id" IS 'Llave foranea logica a la persona que recibe la notificacion, registrada en terceros';

COMMENT ON COLUMN "cotizaciones"."notificacion"."canal_id" IS 'Llave foranea logica al catalogo de canales de notificacion que estara en Parametros';

COMMENT ON COLUMN "cotizaciones"."notificacion"."estado_envio_id" IS 'Llave foranea logica al catalogo de estados de envio que estara en Parametros. Resultado del envio';

COMMENT ON COLUMN "cotizaciones"."notificacion"."intentos" IS 'Cantidad de intentos de envio realizados';

COMMENT ON COLUMN "cotizaciones"."notificacion"."ultimo_error" IS 'Mensaje del ultimo error de envio. Vacio si no hubo error';

COMMENT ON COLUMN "cotizaciones"."notificacion"."fecha_envio" IS 'Fecha y hora en que la notificacion se envio con exito';

COMMENT ON COLUMN "cotizaciones"."notificacion"."fecha_modificacion" IS 'Fecha y hora de la ultima modificacion del registro';

ALTER TABLE "cotizaciones"."solicitud_necesidad" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_item" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_item" ADD FOREIGN KEY ("cargue_masivo_id") REFERENCES "cotizaciones"."cargue_masivo" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cargue_masivo" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cargue_masivo_error" ADD FOREIGN KEY ("cargue_masivo_id") REFERENCES "cotizaciones"."cargue_masivo" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_pago_parcial" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_ciiu" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."documento_solicitud" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_requisito_documental" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_proveedor" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."observacion" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."observacion" ADD FOREIGN KEY ("proveedor_convocado_id") REFERENCES "cotizaciones"."solicitud_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."respuesta_observacion" ADD FOREIGN KEY ("observacion_id") REFERENCES "cotizaciones"."observacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion" ADD FOREIGN KEY ("proveedor_convocado_id") REFERENCES "cotizaciones"."solicitud_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_version" ADD FOREIGN KEY ("cotizacion_id") REFERENCES "cotizaciones"."cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_version" ADD FOREIGN KEY ("solicitud_modificacion_id") REFERENCES "cotizaciones"."solicitud_modificacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_version" ADD FOREIGN KEY ("periodo_subsanacion_id") REFERENCES "cotizaciones"."periodo_subsanacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_version_item" ADD FOREIGN KEY ("cotizacion_version_id") REFERENCES "cotizaciones"."cotizacion_version" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_version_item" ADD FOREIGN KEY ("solicitud_item_id") REFERENCES "cotizaciones"."solicitud_item" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."documento_cotizacion" ADD FOREIGN KEY ("cotizacion_id") REFERENCES "cotizaciones"."cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."documento_cotizacion" ADD FOREIGN KEY ("requisito_documental_id") REFERENCES "cotizaciones"."solicitud_requisito_documental" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."documento_cotizacion" ADD FOREIGN KEY ("cotizacion_version_id") REFERENCES "cotizaciones"."cotizacion_version" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."documento_cotizacion" ADD FOREIGN KEY ("reemplaza_a_id") REFERENCES "cotizaciones"."documento_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."periodo_subsanacion" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."periodo_subsanacion" ADD FOREIGN KEY ("proveedor_convocado_id") REFERENCES "cotizaciones"."solicitud_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."decision_solicitante" ADD FOREIGN KEY ("cotizacion_id") REFERENCES "cotizaciones"."cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."decision_solicitante" ADD FOREIGN KEY ("cotizacion_version_id") REFERENCES "cotizaciones"."cotizacion_version" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."validacion_ordenador" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."validacion_ordenador" ADD FOREIGN KEY ("decision_cotizacion_id") REFERENCES "cotizaciones"."decision_solicitante" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_modificacion" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_modificacion_estado" ADD FOREIGN KEY ("solicitud_modificacion_id") REFERENCES "cotizaciones"."solicitud_modificacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_modificacion_detalle" ADD FOREIGN KEY ("solicitud_modificacion_id") REFERENCES "cotizaciones"."solicitud_modificacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_modificacion_detalle" ADD FOREIGN KEY ("solicitud_item_id") REFERENCES "cotizaciones"."solicitud_item" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_estado" ADD FOREIGN KEY ("solicitud_id") REFERENCES "cotizaciones"."solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cotizacion_estado" ADD FOREIGN KEY ("cotizacion_id") REFERENCES "cotizaciones"."cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_proveedor" ADD FOREIGN KEY ("fecha_creacion") REFERENCES "cotizaciones"."solicitud_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;
