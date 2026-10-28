BEGIN;

-- Solicitudes de cotización
INSERT INTO cotizaciones.solicitud_cotizacion
  (numero_solicitud, consecutivo, vigencia, unidad_ejecutora_id, tipo_proceso_id, estado_id, titulo, objetivo,
   requisitos, observaciones, plazo_ejecucion, criterio_seleccion, plan_accion_id, dependencia_solicitante_id,
   dependencia_destino_id, ordenador_gasto_id, responsable_tercero_id, tipo_necesidad_id, tipo_contrato_id,
   medio_pago_id, tipo_forma_pago_id, area_conocimiento_id, nucleo_basico_id, fecha_apertura, fecha_cierre,
   fecha_procesamiento)
VALUES
  ('SC-PRUEBA-001', 1, 2026, 1, 1, 4, '[PRUEBA] Equipos de computo',
   'Dotar los laboratorios de la Facultad de Ingeniería con equipos de cómputo actualizados.',
   'Garantía mínima de 3 años y soporte técnico en sitio.', 'Entrega en la sede principal.',
   '60 días calendario', 'Menor precio', 1, 10, 10, 100, 200, 1, 1, 1, 2, 1, 1,
   '2026-11-03', '2026-11-13', '2026-11-03 08:00:00'),
  (NULL, NULL, 2026, 1, 1, 1, '[PRUEBA] Mantenimiento laboratorio',
   'Mantenimiento preventivo de los equipos del laboratorio de química.',
   'Técnicos certificados.', NULL, '30 días calendario', 'Menor precio', NULL, 20, 20, 100, 201, 2, 2, 1, 1,
   NULL, NULL, NULL, NULL, NULL),
  ('SC-PRUEBA-003', 2, 2025, 1, 1, 6, '[PRUEBA] Material bibliografico',
   'Adquisición de material bibliográfico especializado para la biblioteca.',
   'Ediciones originales.', NULL, '45 días calendario', 'Mejor relación costo beneficio', 2, 10, 30, 100, 200,
   1, 1, 1, 2, 2, 3, '2025-08-04', '2025-08-11', '2025-08-04 09:30:00');

-- Necesidades (únicas por solicitud y número de necesidad)
INSERT INTO cotizaciones.solicitud_necesidad
  (solicitud_id, numero_necesidad, numero_disponibilidad, valor_disponibilidad, valor_cdp, rubro_codigo, rubro_nombre,
   dependencia_necesidad, estado_necesidad)
SELECT s.id, v.numero_necesidad, v.numero_disponibilidad, v.valor, v.valor, v.rubro_codigo, v.rubro_nombre,
       v.dependencia, v.estado
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        '2026-0101', 'CDP-2026-0501', 85000000.0, '3-01-001-02', 'Equipos de cómputo', 'Facultad de Ingeniería', 'APROBADA'),
  ('[PRUEBA] Equipos de computo',        '2026-0102', 'CDP-2026-0502', 15000000.0, '3-01-001-02', 'Equipos de cómputo', 'Facultad de Ingeniería', 'APROBADA'),
  ('[PRUEBA] Mantenimiento laboratorio', '2026-0230', NULL,            12000000.0, '3-02-002-05', 'Mantenimiento',      'Laboratorio de Química', 'EN TRAMITE'),
  ('[PRUEBA] Material bibliografico',    '2025-0877', 'CDP-2025-0311', 23500000.0, '3-01-004-01', 'Material bibliográfico', 'Biblioteca', 'APROBADA')
) AS v(titulo, numero_necesidad, numero_disponibilidad, valor, rubro_codigo, rubro_nombre, dependencia, estado)
  ON s.titulo = v.titulo;

-- Cargue masivo de ítems (primera solicitud) y su error
INSERT INTO cotizaciones.cargue_masivo
  (solicitud_id, plantilla_id, documento_id, nombre_archivo, cargado_por_tercero_id, total_registros, aceptados,
   rechazados, estado_id)
SELECT s.id, 1, 5001, 'PRUEBA_items_equipos.xlsx', 200, 3, 2, 1, 2
FROM cotizaciones.solicitud_cotizacion s
WHERE s.titulo = '[PRUEBA] Equipos de computo';

INSERT INTO cotizaciones.cargue_masivo_error (cargue_masivo_id, hoja, fila, campo, valor_detectado, motivo_id, detalle)
SELECT c.id, 'Items', 4, 'Cantidad', 'diez', 1, 'La cantidad debe ser numérica'
FROM cotizaciones.cargue_masivo c
WHERE c.nombre_archivo = 'PRUEBA_items_equipos.xlsx';

-- Ítems: dos del cargue masivo, uno manual y uno manual inactivo en la primera solicitud
INSERT INTO cotizaciones.solicitud_item
  (solicitud_id, tipo_bien_servicio_id, nombre, descripcion, unidad_medida_id, cantidad, cargue_masivo_id, codigo_unspsc)
SELECT s.id, v.tipo, v.nombre, v.descripcion, v.unidad, v.cantidad,
       (SELECT c.id FROM cotizaciones.cargue_masivo c WHERE c.nombre_archivo = v.archivo), v.unspsc
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        1, 'Computador de escritorio', 'Procesador de 8 núcleos, 16 GB de RAM, SSD de 512 GB, monitor de 24 pulgadas', 1, 25.00, 'PRUEBA_items_equipos.xlsx', '43211507'),
  ('[PRUEBA] Equipos de computo',        1, 'Portátil',                 'Procesador de 8 núcleos, 16 GB de RAM, pantalla de 14 pulgadas',             1, 10.00, 'PRUEBA_items_equipos.xlsx', '43211503'),
  ('[PRUEBA] Equipos de computo',        1, 'Regulador de voltaje',     'Regulador de 1000 VA con 8 tomas',                                           1,  5.00, NULL,                        '39121006'),
  ('[PRUEBA] Equipos de computo',        1, 'Impresora láser',          'Impresora láser monocromática',                                              1,  2.00, NULL,                        '43212105'),
  ('[PRUEBA] Mantenimiento laboratorio', 2, 'Mantenimiento preventivo', 'Mantenimiento preventivo semestral de campanas extractoras',                NULL, 1.00, NULL,                        NULL),
  ('[PRUEBA] Mantenimiento laboratorio', 2, 'Calibración de balanzas',  'Calibración de balanzas analíticas con certificado',                       NULL, 1.00, NULL,                        NULL),
  ('[PRUEBA] Material bibliografico',    1, 'Libro de cálculo',         'Cálculo de una variable, 9.a edición',                                       1, 40.00, NULL,                        '55101509')
) AS v(titulo, tipo, nombre, descripcion, unidad, cantidad, archivo, unspsc)
  ON s.titulo = v.titulo;

UPDATE cotizaciones.solicitud_item SET activo = false
WHERE nombre = 'Impresora láser'
  AND solicitud_id = (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo = '[PRUEBA] Equipos de computo');

-- Pagos parciales (los porcentajes de cada solicitud suman 1)
INSERT INTO cotizaciones.solicitud_pago_parcial (solicitud_id, condicion, porcentaje)
SELECT s.id, v.condicion, v.porcentaje
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',     'Contra entrega del 40 % de los equipos', 0.4000),
  ('[PRUEBA] Equipos de computo',     'Contra entrega e instalación del 60 % restante', 0.6000),
  ('[PRUEBA] Material bibliografico', 'Contra entrega total del material', 1.0000)
) AS v(titulo, condicion, porcentaje)
  ON s.titulo = v.titulo;

-- Actividades económicas CIIU (únicas por solicitud y actividad)
INSERT INTO cotizaciones.solicitud_ciiu (solicitud_id, actividad_economica_id)
SELECT s.id, v.actividad
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        4651),
  ('[PRUEBA] Equipos de computo',        4741),
  ('[PRUEBA] Mantenimiento laboratorio', 3312),
  ('[PRUEBA] Material bibliografico',    4761)
) AS v(titulo, actividad)
  ON s.titulo = v.titulo;

-- Documentos de la solicitud
INSERT INTO cotizaciones.documento_solicitud
  (solicitud_id, tipo_documento_id, documento_id, nombre_archivo, extension, tamano_bytes, cargado_por_tercero_id)
SELECT s.id, v.tipo, v.documento, v.archivo, v.extension, v.tamano, v.tercero
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        1, 5002, 'estudios_previos.pdf',   'pdf',  245760, 200),
  ('[PRUEBA] Equipos de computo',        2, 5003, 'anexo_tecnico.pdf',      'pdf',  102400, 200),
  ('[PRUEBA] Mantenimiento laboratorio', 1, 5004, 'estudios_previos.docx',  'docx',  51200, 201)
) AS v(titulo, tipo, documento, archivo, extension, tamano, tercero)
  ON s.titulo = v.titulo;

-- Requisitos documentales (únicos por solicitud y tipo de documento)
INSERT INTO cotizaciones.solicitud_requisito_documental (solicitud_id, tipo_documento_id, clasificacion_id, instruccion)
SELECT s.id, v.tipo, v.clasificacion, v.instruccion
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        10, 1, 'Adjuntar RUT actualizado'),
  ('[PRUEBA] Equipos de computo',        11, 2, 'Adjuntar certificación del fabricante, si la tiene'),
  ('[PRUEBA] Mantenimiento laboratorio', 10, 1, 'Adjuntar RUT actualizado'),
  ('[PRUEBA] Material bibliografico',    10, 1, 'Adjuntar RUT actualizado')
) AS v(titulo, tipo, clasificacion, instruccion)
  ON s.titulo = v.titulo;

-- Proveedores convocados (únicos por solicitud y proveedor); el 103 de la primera solicitud queda inactivo
INSERT INTO cotizaciones.solicitud_proveedor (solicitud_id, proveedor_id, fecha_convocatoria)
SELECT s.id, v.proveedor, v.fecha::timestamp
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',     101, '2026-11-03 08:00:00'),
  ('[PRUEBA] Equipos de computo',     102, '2026-11-03 08:00:00'),
  ('[PRUEBA] Equipos de computo',     103, '2026-11-03 08:00:00'),
  ('[PRUEBA] Material bibliografico', 101, '2025-08-04 09:30:00'),
  ('[PRUEBA] Material bibliografico', 104, '2025-08-04 09:30:00')
) AS v(titulo, proveedor, fecha)
  ON s.titulo = v.titulo;

UPDATE cotizaciones.solicitud_proveedor SET activo = false
WHERE proveedor_id = 103
  AND solicitud_id = (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo = '[PRUEBA] Equipos de computo');

-- Historial de estados de cada solicitud
INSERT INTO cotizaciones.solicitud_estado (solicitud_id, estado_id, responsable_tercero_id, justificacion, fecha_creacion)
SELECT s.id, v.estado, v.responsable, v.justificacion, v.fecha::timestamp
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo',        1, 200,  'Creación de la solicitud',                        '2026-10-28 10:00:00'),
  ('[PRUEBA] Equipos de computo',        2, NULL, 'Envío a proveedores con validación exitosa',      '2026-11-03 08:00:00'),
  ('[PRUEBA] Equipos de computo',        3, 200,  'Solicitud de modificación 1: ajuste de cantidades', '2026-11-05 09:00:00'),
  ('[PRUEBA] Equipos de computo',        4, 100,  'Modificación 1 aprobada por el ordenador del gasto', '2026-11-05 15:00:00'),
  ('[PRUEBA] Mantenimiento laboratorio', 1, 201,  'Creación de la solicitud',                        '2026-10-30 11:00:00'),
  ('[PRUEBA] Material bibliografico',    1, 200,  'Creación de la solicitud',                        '2025-07-28 10:00:00'),
  ('[PRUEBA] Material bibliografico',    2, NULL, 'Envío a proveedores con validación exitosa',      '2025-08-04 09:30:00'),
  ('[PRUEBA] Material bibliografico',    5, NULL, 'Cierre automático por fecha límite',              '2025-08-11 23:59:00'),
  ('[PRUEBA] Material bibliografico',    6, 100,  'Adjudicada al proveedor 104',                     '2025-08-20 16:00:00')
) AS v(titulo, estado, responsable, justificacion, fecha)
  ON s.titulo = v.titulo;

-- Modificaciones de la primera solicitud (únicas por solicitud y número):
-- la 1 fue aprobada y la 2 no fue aprobada
INSERT INTO cotizaciones.solicitud_modificacion (solicitud_id, numero, fecha_creacion)
SELECT s.id, v.numero, v.fecha::timestamp
FROM cotizaciones.solicitud_cotizacion s
JOIN (VALUES
  ('[PRUEBA] Equipos de computo', 1, '2026-11-05 09:00:00'),
  ('[PRUEBA] Equipos de computo', 2, '2026-11-07 10:00:00')
) AS v(titulo, numero, fecha)
  ON s.titulo = v.titulo;

INSERT INTO cotizaciones.solicitud_modificacion_estado
  (solicitud_modificacion_id, estado_id, responsable_tercero_id, justificacion, fecha_creacion)
SELECT m.id, v.estado, v.responsable, v.justificacion, v.fecha::timestamp
FROM cotizaciones.solicitud_modificacion m
JOIN cotizaciones.solicitud_cotizacion s ON s.id = m.solicitud_id AND s.titulo = '[PRUEBA] Equipos de computo'
JOIN (VALUES
  (1, 1, 200, 'Ajuste de cantidades y ampliación de plazo', '2026-11-05 09:00:00'),
  (1, 2, 100, 'Se aprueba la modificación',                 '2026-11-05 15:00:00'),
  (2, 1, 200, 'Cambio de especificación del portátil',      '2026-11-07 10:00:00'),
  (2, 3, 100, 'El cambio altera el objeto de la solicitud', '2026-11-07 16:00:00')
) AS v(numero, estado, responsable, justificacion, fecha)
  ON m.numero = v.numero;

-- Detalle de los cambios: valores anterior y nuevo en JSON; el cambio de plazo no tiene ítem
INSERT INTO cotizaciones.solicitud_modificacion_detalle
  (solicitud_modificacion_id, tipo_cambio_id, solicitud_item_id, valor_anterior, valor_nuevo, justificacion)
SELECT m.id, v.tipo_cambio,
       (SELECT i.id FROM cotizaciones.solicitud_item i WHERE i.solicitud_id = s.id AND i.nombre = v.item),
       v.anterior::jsonb, v.nuevo::jsonb, v.justificacion
FROM cotizaciones.solicitud_modificacion m
JOIN cotizaciones.solicitud_cotizacion s ON s.id = m.solicitud_id AND s.titulo = '[PRUEBA] Equipos de computo'
JOIN (VALUES
  (1, 1, 'Computador de escritorio', '{"Cantidad": 20}',                '{"Cantidad": 25}',                'Se suman cinco equipos para la sala de posgrado'),
  (1, 2, NULL,                       '{"FechaCierre": "2026-11-10"}',  '{"FechaCierre": "2026-11-13"}',  'Ampliación de plazo por el cambio de cantidades'),
  (2, 1, 'Portátil',                 '{"Descripcion": "Pantalla de 14 pulgadas"}', '{"Descripcion": "Pantalla de 16 pulgadas"}', 'Solicitud del área usuaria')
) AS v(numero, tipo_cambio, item, anterior, nuevo, justificacion)
  ON m.numero = v.numero;

-- Observaciones del foro (primera solicitud) y respuesta pública
INSERT INTO cotizaciones.observacion (solicitud_id, proveedor_convocado_id, texto, estado_id, fecha_vista)
SELECT s.id, p.id, v.texto, v.estado, v.vista::timestamp
FROM cotizaciones.solicitud_cotizacion s
JOIN cotizaciones.solicitud_proveedor p ON p.solicitud_id = s.id
JOIN (VALUES
  (101, '¿Se aceptan equipos con procesador de 10 núcleos?', 3, '2026-11-04 10:00:00'),
  (102, '¿La garantía de 3 años debe ser directa del fabricante?', 1, NULL)
) AS v(proveedor, texto, estado, vista)
  ON p.proveedor_id = v.proveedor
WHERE s.titulo = '[PRUEBA] Equipos de computo';

INSERT INTO cotizaciones.respuesta_observacion (observacion_id, autor_tercero_id, texto, publica)
SELECT o.id, 200, 'Sí, se aceptan procesadores iguales o superiores a 8 núcleos.', true
FROM cotizaciones.observacion o
JOIN cotizaciones.solicitud_cotizacion s ON s.id = o.solicitud_id AND s.titulo = '[PRUEBA] Equipos de computo'
WHERE o.texto = '¿Se aceptan equipos con procesador de 10 núcleos?';

COMMIT;

-- Limpieza: elimina todos los datos de prueba (descomentar y ejecutar)
--
-- BEGIN;
-- DELETE FROM cotizaciones.respuesta_observacion WHERE observacion_id IN
--   (SELECT o.id FROM cotizaciones.observacion o JOIN cotizaciones.solicitud_cotizacion s ON s.id = o.solicitud_id WHERE s.titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.observacion WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_modificacion_detalle WHERE solicitud_modificacion_id IN
--   (SELECT m.id FROM cotizaciones.solicitud_modificacion m JOIN cotizaciones.solicitud_cotizacion s ON s.id = m.solicitud_id WHERE s.titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_modificacion_estado WHERE solicitud_modificacion_id IN
--   (SELECT m.id FROM cotizaciones.solicitud_modificacion m JOIN cotizaciones.solicitud_cotizacion s ON s.id = m.solicitud_id WHERE s.titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_modificacion WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_estado WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.cargue_masivo_error WHERE cargue_masivo_id IN
--   (SELECT c.id FROM cotizaciones.cargue_masivo c JOIN cotizaciones.solicitud_cotizacion s ON s.id = c.solicitud_id WHERE s.titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_item WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.cargue_masivo WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_proveedor WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_requisito_documental WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.documento_solicitud WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_ciiu WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_pago_parcial WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_necesidad WHERE solicitud_id IN (SELECT id FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%');
-- DELETE FROM cotizaciones.solicitud_cotizacion WHERE titulo LIKE '[PRUEBA]%';
-- COMMIT;
