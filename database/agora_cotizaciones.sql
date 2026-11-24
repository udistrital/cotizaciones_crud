CREATE SCHEMA "cotizaciones";

CREATE TABLE "cotizaciones"."forma_pago" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "tipo_forma_pago_id" integer NOT NULL,
  "condicion" text NOT NULL,
  "porcentaje_pago" numeric(5,4) NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."hist_det_solicitud_cotizacion" (
  "id" serial PRIMARY KEY,
  "det_solicitud_cotizacion_id" integer NOT NULL,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "nombre" varchar(300) NOT NULL,
  "descripcion" text NOT NULL,
  "tipo_bien_servicio_id" integer NOT NULL,
  "unidad_medida_id" integer NOT NULL,
  "cantidad" numeric(14,2) NOT NULL,
  "registro_actual" boolean NOT NULL,
  "tipo_registro" text NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."hist_det_cotizacion" (
  "id" serial PRIMARY KEY,
  "det_cotizacion_id" integer NOT NULL,
  "mst_cotizacion_id" integer NOT NULL,
  "det_solicitud_cotizacion_id" integer NOT NULL,
  "valor_unitario" numeric(20,7),
  "iva" numeric(5,4) NOT NULL,
  "ficha_tecnica" text NOT NULL,
  "registro_actual" boolean NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."hist_mst_cotizacion" (
  "id" serial PRIMARY KEY,
  "mst_cotizacion_id" integer,
  "fecha_envio" timestamp NOT NULL,
  "descripcion" text,
  "informacion_entrega" varchar(500),
  "descuentos" text,
  "observaciones" text,
  "fecha_vencimiento" timestamp NOT NULL,
  "soporte_cotizacion" uuid NOT NULL,
  "registro_actual" boolean NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."hist_mst_solicitud_cotizacion" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "numero_solicitud" varchar(20),
  "vigencia" smallint,
  "consecutivo" integer,
  "fecha_procesamiento" timestamp NOT NULL,
  "estado_id" integer NOT NULL,
  "unidad_ejecutora_id" integer NOT NULL,
  "tipo_necesidad_id" integer NOT NULL,
  "titulo" text NOT NULL,
  "fecha_apertura" date NOT NULL,
  "fecha_cierre" date NOT NULL,
  "objetivo" text NOT NULL,
  "requisitos" text NOT NULL,
  "observaciones" text NOT NULL,
  "ordenador_gasto_id" integer NOT NULL,
  "dependencia_solicitante_id" integer NOT NULL,
  "justificacion_seleccion" text NOT NULL,
  "medio_pago_id" integer NOT NULL,
  "plan_accion_id" integer,
  "responsable_tercero_id" integer NOT NULL,
  "metodo_seleccion_id" integer,
  "numero_necesidad" numeric(15,0),
  "dependencia_destino_id" integer NOT NULL,
  "tipo_contrato_id" integer,
  "criterio_seleccion" varchar(250),
  "plazo_ejecucion" varchar(250),
  "registro_actual" boolean NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."rel_solicitud_cotizacion_proveedor" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "proveedor_id" integer NOT NULL,
  "notificacion_id" integer,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."mst_solicitud_cotizacion" (
  "id" serial PRIMARY KEY,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."mst_cotizacion" (
  "id" serial PRIMARY KEY,
  "rel_solicitud_cotizacion_proveedor_id" integer,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" (
  "id" serial PRIMARY KEY,
  "responsable_tercero_id" integer NOT NULL,
  "justificacion" text NOT NULL,
  "estado_hist_id" integer NOT NULL,
  "hist_mst_solicitud_cotizacion_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."solicitud_cotizacion_actividad_economica" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "actividad_economica_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "enlace" uuid NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."solicitud_cotizacion_nucleo_basico" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "nucleo_basico_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."solicitud_cotizacion_observacion_proveedor" (
  "id" serial PRIMARY KEY,
  "texto" text,
  "rel_solicitud_cotizacion_proveedor_id" integer NOT NULL,
  "visto_solicitante" boolean NOT NULL,
  "fecha_visto_solicitante" timestamp,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."decision_solicitante" (
  "id" serial PRIMARY KEY,
  "rel_solicitud_cotizacion_proveedor_id" integer NOT NULL,
  "mensaje_a_proveedor" text NOT NULL,
  "resultado_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."validacion_ordenador" (
  "id" serial PRIMARY KEY,
  "mst_solicitud_cotizacion_id" integer NOT NULL,
  "observaciones" text NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."cargue_masivo" (
  "id" serial PRIMARY KEY,
  "hist_mst_solicitud_cotizacion_id" integer NOT NULL,
  "enlace_plantilla" text,
  "cargado_por_tercero_id" integer NOT NULL,
  "total_registros" integer,
  "aceptados" integer,
  "rechazados" integer,
  "respuesta" json,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."notificacion" (
  "id" serial PRIMARY KEY,
  "proveedor_notificado_id" integer NOT NULL,
  "tercero_notifica_id" integer NOT NULL,
  "estado_envio_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" (
  "id" serial PRIMARY KEY,
  "observacion_id" integer NOT NULL,
  "texto" text NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" (
  "id" serial PRIMARY KEY,
  "maestro_id" integer NOT NULL,
  "detalle_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE TABLE "cotizaciones"."rel_hist_mst_det_cotizacion" (
  "id" serial PRIMARY KEY,
  "maestro_id" integer NOT NULL,
  "detalle_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT (true),
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "usuario_creacion_id" integer,
  "fecha_modificacion" timestamp,
  "usuario_modificacion_id" integer
);

CREATE INDEX "pk_forma_pago" ON "cotizaciones"."forma_pago" ("id");

CREATE INDEX "idx_forma_pago_mst_solicitud_cotizacion_id" ON "cotizaciones"."forma_pago" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_forma_pago_tipo_forma_pago_id" ON "cotizaciones"."forma_pago" ("tipo_forma_pago_id");

CREATE INDEX "idx_forma_pago_activo" ON "cotizaciones"."forma_pago" ("activo");

CREATE INDEX "idx_forma_pago_fecha_creacion" ON "cotizaciones"."forma_pago" ("fecha_creacion");

CREATE INDEX "idx_forma_pago_usuario_creacion_id" ON "cotizaciones"."forma_pago" ("usuario_creacion_id");

CREATE INDEX "idx_forma_pago_fecha_modificacion" ON "cotizaciones"."forma_pago" ("fecha_modificacion");

CREATE INDEX "idx_forma_pago_usuario_modificacion_id" ON "cotizaciones"."forma_pago" ("usuario_modificacion_id");

CREATE INDEX "pk_hist_det_solicitud_cotizacion" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("id");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_mst_solicitud_cotizacion_id" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_tipo_bien_servicio_id" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("tipo_bien_servicio_id");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_unidad_medida_id" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("unidad_medida_id");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_activo" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("activo");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_fecha_creacion" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_usuario_creacion_id" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_fecha_modificacion" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_hist_det_solicitud_cotizacion_usuario_modificacion_id" ON "cotizaciones"."hist_det_solicitud_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_hist_det_cotizacion" ON "cotizaciones"."hist_det_cotizacion" ("id");

CREATE INDEX "idx_hist_det_cotizacion_mst_cotizacion_id" ON "cotizaciones"."hist_det_cotizacion" ("mst_cotizacion_id");

CREATE INDEX "idx_hist_det_cotizacion_activo" ON "cotizaciones"."hist_det_cotizacion" ("activo");

CREATE INDEX "idx_hist_det_cotizacion_fecha_creacion" ON "cotizaciones"."hist_det_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_hist_det_cotizacion_usuario_creacion_id" ON "cotizaciones"."hist_det_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_hist_det_cotizacion_fecha_modificacion" ON "cotizaciones"."hist_det_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_hist_det_cotizacion_usuario_modificacion_id" ON "cotizaciones"."hist_det_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_hist_mst_cotizacion" ON "cotizaciones"."hist_mst_cotizacion" ("id");

CREATE INDEX "idx_hist_mst_cotizacion_mst_cotizacion_id" ON "cotizaciones"."hist_mst_cotizacion" ("mst_cotizacion_id");

CREATE INDEX "idx_hist_mst_cotizacion_activo" ON "cotizaciones"."hist_mst_cotizacion" ("activo");

CREATE INDEX "idx_hist_mst_cotizacion_fecha_creacion" ON "cotizaciones"."hist_mst_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_hist_mst_cotizacion_usuario_creacion_id" ON "cotizaciones"."hist_mst_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_hist_mst_cotizacion_fecha_modificacion" ON "cotizaciones"."hist_mst_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_hist_mst_cotizacion_usuario_modificacion_id" ON "cotizaciones"."hist_mst_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_hist_mst_solicitud_cotizacion" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_mst_solicitud_cotizacion_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_estado_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("estado_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_unidad_ejecutora_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("unidad_ejecutora_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_medio_pago_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("medio_pago_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_responsable_tercero_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("responsable_tercero_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_metodo_seleccion_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("metodo_seleccion_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_tipo_contrato_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("tipo_contrato_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_activo" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("activo");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_fecha_creacion" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_usuario_creacion_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_fecha_modificacion" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_usuario_modificacion_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_rel_solicitud_cotizacion_proveedor" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("id");

CREATE INDEX "idx_rl_slctd_ctzcn_prvdr_mst_slctd_ctzcn_id" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_proveedor_id" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("proveedor_id");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_notificacion_id" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("notificacion_id");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_activo" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("activo");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_fecha_creacion" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("fecha_creacion");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_usuario_creacion_id" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("usuario_creacion_id");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_fecha_modificacion" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_rel_solicitud_cotizacion_proveedor_usuario_modificacion_id" ON "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("usuario_modificacion_id");

CREATE INDEX "pk_mst_solicitud_cotizacion" ON "cotizaciones"."mst_solicitud_cotizacion" ("id");

CREATE INDEX "idx_mst_solicitud_cotizacion_activo" ON "cotizaciones"."mst_solicitud_cotizacion" ("activo");

CREATE INDEX "idx_mst_solicitud_cotizacion_fecha_creacion" ON "cotizaciones"."mst_solicitud_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_mst_solicitud_cotizacion_usuario_creacion_id" ON "cotizaciones"."mst_solicitud_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_mst_solicitud_cotizacion_fecha_modificacion" ON "cotizaciones"."mst_solicitud_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_mst_solicitud_cotizacion_usuario_modificacion_id" ON "cotizaciones"."mst_solicitud_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_mst_cotizacion" ON "cotizaciones"."mst_cotizacion" ("id");

CREATE INDEX "idx_mst_cotizacion_rel_solicitud_cotizacion_proveedor_id" ON "cotizaciones"."mst_cotizacion" ("rel_solicitud_cotizacion_proveedor_id");

CREATE INDEX "idx_mst_cotizacion_activo" ON "cotizaciones"."mst_cotizacion" ("activo");

CREATE INDEX "idx_mst_cotizacion_fecha_creacion" ON "cotizaciones"."mst_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_mst_cotizacion_usuario_creacion_id" ON "cotizaciones"."mst_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_mst_cotizacion_fecha_modificacion" ON "cotizaciones"."mst_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_mst_cotizacion_usuario_modificacion_id" ON "cotizaciones"."mst_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_hist_mst_solicitud_cotizacion_estado_hist" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("id");

CREATE INDEX "idx_hst_mst_slctd_ctzcn_estd_hst_rspnsbl_trcr_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("responsable_tercero_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_estado_hist_estado_hist_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("estado_hist_id");

CREATE INDEX "idx_hst_mst_slctd_ctzcn_estd_hst_hst_mst_slctd_ctzcn_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("hist_mst_solicitud_cotizacion_id");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_estado_hist_activo" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("activo");

CREATE INDEX "idx_hist_mst_solicitud_cotizacion_estado_hist_fecha_creacion" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("fecha_creacion");

CREATE INDEX "idx_hst_mst_slctd_ctzcn_estd_hst_usr_crcn_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("usuario_creacion_id");

CREATE INDEX "idx_hst_mst_slctd_ctzcn_estd_hst_fch_mdfccn" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("fecha_modificacion");

CREATE INDEX "idx_hst_mst_slctd_ctzcn_estd_hst_usr_mdfccn_id" ON "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ("usuario_modificacion_id");

CREATE INDEX "pk_solicitud_cotizacion_actividad_economica" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("id");

CREATE INDEX "idx_slctd_ctzcn_actvdd_ecnmc_mst_slctd_ctzcn_id" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_slctd_ctzcn_actvdd_ecnmc_actvdd_ecnmc_id" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("actividad_economica_id");

CREATE INDEX "idx_solicitud_cotizacion_actividad_economica_activo" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("activo");

CREATE INDEX "idx_solicitud_cotizacion_actividad_economica_fecha_creacion" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("fecha_creacion");

CREATE INDEX "idx_slctd_ctzcn_actvdd_ecnmc_usr_crcn_id" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("usuario_creacion_id");

CREATE INDEX "idx_solicitud_cotizacion_actividad_economica_fecha_modificacion" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("fecha_modificacion");

CREATE INDEX "idx_slctd_ctzcn_actvdd_ecnmc_usr_mdfccn_id" ON "cotizaciones"."solicitud_cotizacion_actividad_economica" ("usuario_modificacion_id");

CREATE INDEX "pk_solicitud_cotizacion_especificacion_tecnica" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("id");

CREATE INDEX "idx_slctd_ctzcn_espcfccn_tcnc_mst_slctd_ctzcn_id" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_solicitud_cotizacion_especificacion_tecnica_activo" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("activo");

CREATE INDEX "idx_solicitud_cotizacion_especificacion_tecnica_fecha_creacion" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("fecha_creacion");

CREATE INDEX "idx_slctd_ctzcn_espcfccn_tcnc_usr_crcn_id" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("usuario_creacion_id");

CREATE INDEX "idx_slctd_ctzcn_espcfccn_tcnc_fch_mdfccn" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("fecha_modificacion");

CREATE INDEX "idx_slctd_ctzcn_espcfccn_tcnc_usr_mdfccn_id" ON "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ("usuario_modificacion_id");

CREATE INDEX "pk_solicitud_cotizacion_nucleo_basico" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("id");

CREATE INDEX "idx_slctd_ctzcn_ncl_bsc_mst_slctd_ctzcn_id" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_nucleo_basico_id" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("nucleo_basico_id");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_activo" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("activo");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_fecha_creacion" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("fecha_creacion");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_usuario_creacion_id" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("usuario_creacion_id");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_fecha_modificacion" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("fecha_modificacion");

CREATE INDEX "idx_solicitud_cotizacion_nucleo_basico_usuario_modificacion_id" ON "cotizaciones"."solicitud_cotizacion_nucleo_basico" ("usuario_modificacion_id");

CREATE INDEX "pk_solicitud_cotizacion_observacion_proveedor" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("id");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rl_slctd_ctzcn_prvdr_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("rel_solicitud_cotizacion_proveedor_id");

CREATE INDEX "idx_solicitud_cotizacion_observacion_proveedor_activo" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("activo");

CREATE INDEX "idx_solicitud_cotizacion_observacion_proveedor_fecha_creacion" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("fecha_creacion");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_usr_crcn_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("usuario_creacion_id");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_fch_mdfccn" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_usr_mdfccn_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("usuario_modificacion_id");

CREATE INDEX "pk_decision_solicitante" ON "cotizaciones"."decision_solicitante" ("id");

CREATE INDEX "idx_decision_solicitante_rel_solicitud_cotizacion_proveedor_id" ON "cotizaciones"."decision_solicitante" ("rel_solicitud_cotizacion_proveedor_id");

CREATE INDEX "idx_decision_solicitante_resultado_id" ON "cotizaciones"."decision_solicitante" ("resultado_id");

CREATE INDEX "idx_decision_solicitante_activo" ON "cotizaciones"."decision_solicitante" ("activo");

CREATE INDEX "idx_decision_solicitante_fecha_creacion" ON "cotizaciones"."decision_solicitante" ("fecha_creacion");

CREATE INDEX "idx_decision_solicitante_usuario_creacion_id" ON "cotizaciones"."decision_solicitante" ("usuario_creacion_id");

CREATE INDEX "idx_decision_solicitante_fecha_modificacion" ON "cotizaciones"."decision_solicitante" ("fecha_modificacion");

CREATE INDEX "idx_decision_solicitante_usuario_modificacion_id" ON "cotizaciones"."decision_solicitante" ("usuario_modificacion_id");

CREATE INDEX "pk_validacion_ordenador" ON "cotizaciones"."validacion_ordenador" ("id");

CREATE INDEX "idx_validacion_ordenador_mst_solicitud_cotizacion_id" ON "cotizaciones"."validacion_ordenador" ("mst_solicitud_cotizacion_id");

CREATE INDEX "idx_validacion_ordenador_activo" ON "cotizaciones"."validacion_ordenador" ("activo");

CREATE INDEX "idx_validacion_ordenador_fecha_creacion" ON "cotizaciones"."validacion_ordenador" ("fecha_creacion");

CREATE INDEX "idx_validacion_ordenador_usuario_creacion_id" ON "cotizaciones"."validacion_ordenador" ("usuario_creacion_id");

CREATE INDEX "idx_validacion_ordenador_fecha_modificacion" ON "cotizaciones"."validacion_ordenador" ("fecha_modificacion");

CREATE INDEX "idx_validacion_ordenador_usuario_modificacion_id" ON "cotizaciones"."validacion_ordenador" ("usuario_modificacion_id");

CREATE INDEX "pk_cargue_masivo" ON "cotizaciones"."cargue_masivo" ("id");

CREATE INDEX "idx_cargue_masivo_hist_mst_solicitud_cotizacion_id" ON "cotizaciones"."cargue_masivo" ("hist_mst_solicitud_cotizacion_id");

CREATE INDEX "idx_cargue_masivo_cargado_por_tercero_id" ON "cotizaciones"."cargue_masivo" ("cargado_por_tercero_id");

CREATE INDEX "idx_cargue_masivo_activo" ON "cotizaciones"."cargue_masivo" ("activo");

CREATE INDEX "idx_cargue_masivo_fecha_creacion" ON "cotizaciones"."cargue_masivo" ("fecha_creacion");

CREATE INDEX "idx_cargue_masivo_usuario_creacion_id" ON "cotizaciones"."cargue_masivo" ("usuario_creacion_id");

CREATE INDEX "idx_cargue_masivo_fecha_modificacion" ON "cotizaciones"."cargue_masivo" ("fecha_modificacion");

CREATE INDEX "idx_cargue_masivo_usuario_modificacion_id" ON "cotizaciones"."cargue_masivo" ("usuario_modificacion_id");

CREATE INDEX "pk_notificacion" ON "cotizaciones"."notificacion" ("id");

CREATE INDEX "idx_notificacion_proveedor_notificado_id" ON "cotizaciones"."notificacion" ("proveedor_notificado_id");

CREATE INDEX "idx_notificacion_tercero_notifica_id" ON "cotizaciones"."notificacion" ("tercero_notifica_id");

CREATE INDEX "idx_notificacion_estado_envio_id" ON "cotizaciones"."notificacion" ("estado_envio_id");

CREATE INDEX "idx_notificacion_activo" ON "cotizaciones"."notificacion" ("activo");

CREATE INDEX "idx_notificacion_fecha_creacion" ON "cotizaciones"."notificacion" ("fecha_creacion");

CREATE INDEX "idx_notificacion_usuario_creacion_id" ON "cotizaciones"."notificacion" ("usuario_creacion_id");

CREATE INDEX "idx_notificacion_fecha_modificacion" ON "cotizaciones"."notificacion" ("fecha_modificacion");

CREATE INDEX "idx_notificacion_usuario_modificacion_id" ON "cotizaciones"."notificacion" ("usuario_modificacion_id");

CREATE INDEX "pk_solicitud_cotizacion_observacion_proveedor_respuesta" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("id");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rspst_obsrvcn_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("observacion_id");

CREATE INDEX "idx_solicitud_cotizacion_observacion_proveedor_respuesta_activo" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("activo");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rspst_fch_crcn" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("fecha_creacion");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rspst_usr_crcn_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("usuario_creacion_id");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rspst_fch_mdfccn" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("fecha_modificacion");

CREATE INDEX "idx_slctd_ctzcn_obsrvcn_prvdr_rspst_usr_mdfccn_id" ON "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ("usuario_modificacion_id");

CREATE INDEX "pk_rel_hist_mst_det_solicitud_cotizacion" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("id");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_maestro_id" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("maestro_id");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_detalle_id" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("detalle_id");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_activo" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("activo");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_fecha_creacion" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_usuario_creacion_id" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_rel_hist_mst_det_solicitud_cotizacion_fecha_modificacion" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_rl_hst_mst_dt_slctd_ctzcn_usr_mdfccn_id" ON "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ("usuario_modificacion_id");

CREATE INDEX "pk_rel_hist_mst_det_cotizacion" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("id");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_maestro_id" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("maestro_id");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_detalle_id" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("detalle_id");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_activo" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("activo");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_fecha_creacion" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("fecha_creacion");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_usuario_creacion_id" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("usuario_creacion_id");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_fecha_modificacion" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("fecha_modificacion");

CREATE INDEX "idx_rel_hist_mst_det_cotizacion_usuario_modificacion_id" ON "cotizaciones"."rel_hist_mst_det_cotizacion" ("usuario_modificacion_id");

COMMENT ON COLUMN "cotizaciones"."forma_pago"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."forma_pago"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_det_solicitud_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_det_solicitud_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_det_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_det_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_solicitud_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_solicitud_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_solicitud_cotizacion_proveedor"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_solicitud_cotizacion_proveedor"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."mst_solicitud_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."mst_solicitud_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."mst_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."mst_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_actividad_economica"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_actividad_economica"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_especificacion_tecnica"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_especificacion_tecnica"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_nucleo_basico"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_nucleo_basico"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_observacion_proveedor"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_observacion_proveedor"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."decision_solicitante"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."validacion_ordenador"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."cargue_masivo"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."notificacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."notificacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_hist_mst_det_cotizacion"."usuario_creacion_id" IS 'Posible referencia externa a Terceros CRUD';

COMMENT ON COLUMN "cotizaciones"."rel_hist_mst_det_cotizacion"."usuario_modificacion_id" IS 'Posible referencia externa a Terceros CRUD';

ALTER TABLE "cotizaciones"."forma_pago" ADD CONSTRAINT "fk_forma_pago_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."hist_det_solicitud_cotizacion" ADD CONSTRAINT "fk_hist_det_solicitud_cotizacion_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."hist_det_cotizacion" ADD CONSTRAINT "fk_hist_det_cotizacion_mst_cotizacion" FOREIGN KEY ("mst_cotizacion_id") REFERENCES "cotizaciones"."mst_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."hist_mst_cotizacion" ADD CONSTRAINT "fk_hist_mst_cotizacion_mst_cotizacion" FOREIGN KEY ("mst_cotizacion_id") REFERENCES "cotizaciones"."mst_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."hist_mst_solicitud_cotizacion" ADD CONSTRAINT "fk_hist_mst_solicitud_cotizacion_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_solicitud_cotizacion_proveedor" ADD CONSTRAINT "fk_rel_solicitud_cotizacion_proveedor_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_solicitud_cotizacion_proveedor" ADD CONSTRAINT "fk_rel_solicitud_cotizacion_proveedor_notificacion" FOREIGN KEY ("notificacion_id") REFERENCES "cotizaciones"."notificacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."mst_cotizacion" ADD CONSTRAINT "fk_mst_cotizacion_rel_solicitud_cotizacion_proveedor" FOREIGN KEY ("rel_solicitud_cotizacion_proveedor_id") REFERENCES "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."hist_mst_solicitud_cotizacion_estado_hist" ADD CONSTRAINT "fk_hst_mst_slctd_ctzcn_estd_hst_hst_mst_slctd_ctzcn" FOREIGN KEY ("hist_mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."hist_mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_cotizacion_actividad_economica" ADD CONSTRAINT "fk_slctd_ctzcn_actvdd_ecnmc_mst_slctd_ctzcn" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_cotizacion_especificacion_tecnica" ADD CONSTRAINT "fk_slctd_ctzcn_espcfccn_tcnc_mst_slctd_ctzcn" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_cotizacion_nucleo_basico" ADD CONSTRAINT "fk_solicitud_cotizacion_nucleo_basico_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ADD CONSTRAINT "fk_slctd_ctzcn_obsrvcn_prvdr_rl_slctd_ctzcn_prvdr" FOREIGN KEY ("rel_solicitud_cotizacion_proveedor_id") REFERENCES "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."decision_solicitante" ADD CONSTRAINT "fk_decision_solicitante_rel_solicitud_cotizacion_proveedor" FOREIGN KEY ("rel_solicitud_cotizacion_proveedor_id") REFERENCES "cotizaciones"."rel_solicitud_cotizacion_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."validacion_ordenador" ADD CONSTRAINT "fk_validacion_ordenador_mst_solicitud_cotizacion" FOREIGN KEY ("mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."cargue_masivo" ADD CONSTRAINT "fk_cargue_masivo_hist_mst_solicitud_cotizacion" FOREIGN KEY ("hist_mst_solicitud_cotizacion_id") REFERENCES "cotizaciones"."hist_mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."solicitud_cotizacion_observacion_proveedor_respuesta" ADD CONSTRAINT "fk_slctd_ctzcn_obsrvcn_prvdr_rspst_slctd_ctzcn_obsrvcn_prvdr" FOREIGN KEY ("observacion_id") REFERENCES "cotizaciones"."solicitud_cotizacion_observacion_proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ADD CONSTRAINT "fk_rl_hst_mst_dt_slctd_ctzcn_hst_mst_slctd_ctzcn" FOREIGN KEY ("maestro_id") REFERENCES "cotizaciones"."hist_mst_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_hist_mst_det_solicitud_cotizacion" ADD CONSTRAINT "fk_rl_hst_mst_dt_slctd_ctzcn_hst_dt_slctd_ctzcn" FOREIGN KEY ("detalle_id") REFERENCES "cotizaciones"."hist_det_solicitud_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_hist_mst_det_cotizacion" ADD CONSTRAINT "fk_rel_hist_mst_det_cotizacion_hist_mst_cotizacion" FOREIGN KEY ("maestro_id") REFERENCES "cotizaciones"."hist_mst_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cotizaciones"."rel_hist_mst_det_cotizacion" ADD CONSTRAINT "fk_rel_hist_mst_det_cotizacion_hist_det_cotizacion" FOREIGN KEY ("detalle_id") REFERENCES "cotizaciones"."hist_det_cotizacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;
