package models

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) Constructor(inombre string, iedad int) {
	p.nombre = inombre
	p.edad = iedad
}

func (p *Persona) GetNombre() string { return p.nombre }
func (p *Persona) GetEdad() int      { return p.edad }
