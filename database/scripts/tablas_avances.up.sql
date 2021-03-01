-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-02-16 16:55:20.141

-- Schema
CREATE SCHEMA avances;
SET search_path TO pg_catalog,public,avances;

-- tables
-- Table: avance_legalizacion
CREATE TABLE avances.avance_legalizacion (
    id serial  NOT NULL,
    solicitud_avance_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT uq UNIQUE (solicitud_avance_id) NOT DEFERRABLE  INITIALLY IMMEDIATE,
    CONSTRAINT avance_legalizacion_pk PRIMARY KEY (id)
);

-- Table: avance_legalizacion_tipo
CREATE TABLE avances.avance_legalizacion_tipo (
    id serial  NOT NULL,
    tipo_avance_legalizacio_id int  NOT NULL,
    avance_legalizacion_id int  NOT NULL,
    vigencia int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT avance_legalizacion_tipo_pk PRIMARY KEY (id)
);

-- Table: especificacion_solicitud_tipo_avance
CREATE TABLE avances.especificacion_solicitud_tipo_avance (
    id serial  NOT NULL,
    especificacion_tipo_avance_id int  NOT NULL,
    solicitud_tipo_avance_id int  NOT NULL,
    descripcion varchar(250)  NULL,
    valor decimal(10,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT especificacion_solicitud_tipo_avance_pk PRIMARY KEY (id)
);

-- Table: especificacion_tipo_avance
CREATE TABLE avances.especificacion_tipo_avance (
    id serial  NOT NULL,
    tipo_avance_id int  NOT NULL,
    especificacion_avance_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT especificacion_tipo_avance_pk PRIMARY KEY (id)
);

-- Table: estado_legalizacion_avance_legalizacion
CREATE TABLE avances.estado_legalizacion_avance_legalizacion (
    id serial  NOT NULL,
    estado_legalizacion_id int  NOT NULL,
    avance_legalizacion_id int  NOT NULL,
    observaciones varchar(250)  NOT NULL,
    responsable int  NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT estado_legalizacion_avance_legalizacion_pk PRIMARY KEY (id)
);

-- Table: informacion_estudiante
CREATE TABLE avances.informacion_estudiante (
    id serial  NOT NULL,
    practica_academica_id int  NOT NULL,
    codigo varchar(20)  NOT NULL,
    dias int  NOT NULL,
    valor_recibido decimal(10,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT informacion_estudiante_pk PRIMARY KEY (id)
);

-- Table: movimiento
CREATE TABLE avances.movimiento (
    id serial  NOT NULL,
    tipo_movimiento_id int  NOT NULL,
    avance_legalizacion_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT movimiento_pk PRIMARY KEY (id)
);

-- Table: norma
CREATE TABLE avances.norma (
    id serial  NOT NULL,
    tipo_avance_id int  NOT NULL,
    vigencia int  NOT NULL,
    link_norma text  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT norma_pk PRIMARY KEY (id)
);

-- Table: orden_pago_avance_legalizacion
CREATE TABLE avances.orden_pago_avance_legalizacion (
    id serial  NOT NULL,
    avance_legalizacion_id int  NOT NULL,
    vigencia int  NOT NULL,
    orden_pago_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT orden_pago_avance_legalizacion_pk PRIMARY KEY (id)
);

-- Table: practica_academica
CREATE TABLE avances.practica_academica (
    id serial  NOT NULL,
    avance_legalizacion_tipo_id int  NOT NULL,
    fecha_practica int  NOT NULL,
    destino_id int  NOT NULL,
    proyecto_curricular_id int  NOT NULL,
    valor_avance decimal(10,2)  NOT NULL,
    valor_reintegro decimal(10,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT practica_academica_pk PRIMARY KEY (id)
);

-- Table: registro_avance_legalizacion
CREATE TABLE avances.registro_avance_legalizacion (
    id serial  NOT NULL,
    avance_legalizacion_tipo_id int  NOT NULL,
    fecha_evento timestamp  NOT NULL,
    tipo_documento_tercero_id int  NOT NULL,
    documento_tercero varchar  NOT NULL,
    numero_factura int  NOT NULL,
    trm_fecha_compra int  NOT NULL,
    valor_total decimal(10,2)  NOT NULL,
    renta_id int  NOT NULL,
    ica_id int  NOT NULL,
    iva_id int  NOT NULL,
    valor_renta decimal(10,2)  NOT NULL,
    valor_ica decimal(10,2)  NOT NULL,
    valor_iva decimal(10,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT registro_avance_legalizacion_pk PRIMARY KEY (id)
);

-- Table: reintegro
CREATE TABLE avances.reintegro (
    id serial  NOT NULL,
    avance_legalizacion_id int  NOT NULL,
    causal_reintegro_id int  NOT NULL,
    area_funcional int  NOT NULL,
    concepto varchar  NOT NULL,
    tipo_acta varchar  NOT NULL,
    numero_oficio int  NOT NULL,
    valor int  NOT NULL,
    vigencia int  NOT NULL,
    ingreso int  NOT NULL,
    justificacion varchar(250)  NOT NULL,
    numero_orden decimal(5,2)  NOT NULL,
    disponible boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT reintegro_pk PRIMARY KEY (id)
);

-- Table: requisito_tipo_avance
CREATE TABLE avances.requisito_tipo_avance (
    id serial  NOT NULL,
    tipo_avance_id int  NOT NULL,
    requisito_avance_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT requisito_tipo_avance_pk PRIMARY KEY (id)
);

-- Table: solicitud_avance
CREATE TABLE avances.solicitud_avance (
    id serial  NOT NULL,
    solicitud_id int  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT solicitud_avance_pk PRIMARY KEY (id)
);

-- Table: solicitud_requisito_tipo_avance
CREATE TABLE avances.solicitud_requisito_tipo_avance (
    id serial  NOT NULL,
    solicitud_tipo_avance_id int  NOT NULL,
    requisito_tipo_avance_id int  NOT NULL,
    observaciones varchar(250)  NULL,
    documento varchar(10)  NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT solicitud_requisito_tipo_avance_pk PRIMARY KEY (id)
);

-- Table: solicitud_tipo_avance
CREATE TABLE avances.solicitud_tipo_avance (
    id serial  NOT NULL,
    solicitud_avance_id int  NOT NULL,
    tipo_avance_id int  NOT NULL,
    descripcion varchar(250)  NOT NULL,
    valor decimal(10,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT solicitud_tipo_avance_pk PRIMARY KEY (id)
);

-- Table: tipo_avance
CREATE TABLE avances.tipo_avance (
    id serial  NOT NULL,
    codigo_abreviacion varchar(20)  NOT NULL,
    area_funcional int  NOT NULL,
    nombre varchar(50)  NOT NULL,
    descripcion varchar(250)  NOT NULL,
    numero_orden decimal(5,2)  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    activo boolean  NOT NULL,
    CONSTRAINT tipo_avance_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: avance_legalizacion_solicitud_avance (table: avance_legalizacion)
ALTER TABLE avances.avance_legalizacion ADD CONSTRAINT avance_legalizacion_solicitud_avance
    FOREIGN KEY (solicitud_avance_id)
    REFERENCES avances.solicitud_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: avance_legalizacion_tipo_avance_legalizacion (table: avance_legalizacion_tipo)
ALTER TABLE avances.avance_legalizacion_tipo ADD CONSTRAINT avance_legalizacion_tipo_avance_legalizacion
    FOREIGN KEY (avance_legalizacion_id)
    REFERENCES avances.avance_legalizacion (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: consolidado_retencion_avance_legalizacion_tipo (table: registro_avance_legalizacion)
ALTER TABLE avances.registro_avance_legalizacion ADD CONSTRAINT consolidado_retencion_avance_legalizacion_tipo
    FOREIGN KEY (avance_legalizacion_tipo_id)
    REFERENCES avances.avance_legalizacion_tipo (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: especificacion_tipo_avance_tipo_avance (table: especificacion_tipo_avance)
ALTER TABLE avances.especificacion_tipo_avance ADD CONSTRAINT especificacion_tipo_avance_tipo_avance
    FOREIGN KEY (tipo_avance_id)
    REFERENCES avances.tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: estado_legalizacion_avance_legalizacion_avance_legalizacion (table: estado_legalizacion_avance_legalizacion)
ALTER TABLE avances.estado_legalizacion_avance_legalizacion ADD CONSTRAINT estado_legalizacion_avance_legalizacion_avance_legalizacion
    FOREIGN KEY (avance_legalizacion_id)
    REFERENCES avances.avance_legalizacion (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: informacion_estudiante_practica_academica (table: informacion_estudiante)
ALTER TABLE avances.informacion_estudiante ADD CONSTRAINT informacion_estudiante_practica_academica
    FOREIGN KEY (practica_academica_id)
    REFERENCES avances.practica_academica (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: movimientos_avance_legalizacion (table: movimiento)
ALTER TABLE avances.movimiento ADD CONSTRAINT movimientos_avance_legalizacion
    FOREIGN KEY (avance_legalizacion_id)
    REFERENCES avances.avance_legalizacion (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: norma_tipo_avance (table: norma)
ALTER TABLE avances.norma ADD CONSTRAINT norma_tipo_avance
    FOREIGN KEY (tipo_avance_id)
    REFERENCES avances.tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: orden_pago_avance_legalizacion_avance_legalizacion (table: orden_pago_avance_legalizacion)
ALTER TABLE avances.orden_pago_avance_legalizacion ADD CONSTRAINT orden_pago_avance_legalizacion_avance_legalizacion
    FOREIGN KEY (avance_legalizacion_id)
    REFERENCES avances.avance_legalizacion (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: practica_academica_avance_legalizacion_tipo (table: practica_academica)
ALTER TABLE avances.practica_academica ADD CONSTRAINT practica_academica_avance_legalizacion_tipo
    FOREIGN KEY (avance_legalizacion_tipo_id)
    REFERENCES avances.avance_legalizacion_tipo (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: reintegro_avance_legalizacion (table: reintegro)
ALTER TABLE avances.reintegro ADD CONSTRAINT reintegro_avance_legalizacion
    FOREIGN KEY (avance_legalizacion_id)
    REFERENCES avances.avance_legalizacion (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: requisito_tipo_avance_tipo_avance (table: requisito_tipo_avance)
ALTER TABLE avances.requisito_tipo_avance ADD CONSTRAINT requisito_tipo_avance_tipo_avance
    FOREIGN KEY (tipo_avance_id)
    REFERENCES avances.tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_especificaciones_requisito_tipo_avance_especificacion_tipo_avance (table: especificacion_solicitud_tipo_avance)
ALTER TABLE avances.especificacion_solicitud_tipo_avance ADD CONSTRAINT solicitud_especificaciones_requisito_tipo_avance_especificacion_tipo_avance
    FOREIGN KEY (especificacion_tipo_avance_id)
    REFERENCES avances.especificacion_tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_especificaciones_requisito_tipo_avance_solicitud_tipo_avance (table: especificacion_solicitud_tipo_avance)
ALTER TABLE avances.especificacion_solicitud_tipo_avance ADD CONSTRAINT solicitud_especificaciones_requisito_tipo_avance_solicitud_tipo_avance
    FOREIGN KEY (solicitud_tipo_avance_id)
    REFERENCES avances.solicitud_tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_requisito_tipo_avance_requisito_tipo_avance (table: solicitud_requisito_tipo_avance)
ALTER TABLE avances.solicitud_requisito_tipo_avance ADD CONSTRAINT solicitud_requisito_tipo_avance_requisito_tipo_avance
    FOREIGN KEY (requisito_tipo_avance_id)
    REFERENCES avances.requisito_tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_requisito_tipo_avance_solicitud_tipo_avance (table: solicitud_requisito_tipo_avance)
ALTER TABLE avances.solicitud_requisito_tipo_avance ADD CONSTRAINT solicitud_requisito_tipo_avance_solicitud_tipo_avance
    FOREIGN KEY (solicitud_tipo_avance_id)
    REFERENCES avances.solicitud_tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_tipo_avance_solicitud_avance (table: solicitud_tipo_avance)
ALTER TABLE avances.solicitud_tipo_avance ADD CONSTRAINT solicitud_tipo_avance_solicitud_avance
    FOREIGN KEY (solicitud_avance_id)
    REFERENCES avances.solicitud_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: solicitud_tipo_avance_tipo_avance (table: solicitud_tipo_avance)
ALTER TABLE avances.solicitud_tipo_avance ADD CONSTRAINT solicitud_tipo_avance_tipo_avance
    FOREIGN KEY (tipo_avance_id)
    REFERENCES avances.tipo_avance (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

