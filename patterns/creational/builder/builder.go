package builder

type BuildProcess interface {
	SetAmmo() BuildProcess
	SetDamage() BuildProcess
	SetStructure() BuildProcess
	GetGun() Gun
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetAmmo().SetDamage().SetStructure()
}

type Gun struct {
	Ammo      int
	Damage    int
	Structure string
}

type Pistol struct {
	g Gun
}

func (p *Pistol) SetAmmo() BuildProcess {
	p.g.Ammo = 10
	return p
}

func (p *Pistol) SetDamage() BuildProcess {
	p.g.Damage = 5
	return p
}

func (p *Pistol) SetStructure() BuildProcess {
	p.g.Structure = "Pistol"
	return p
}

func (p *Pistol) GetGun() Gun {
	return p.g
}

type Rifle struct {
	g Gun
}

func (r *Rifle) SetAmmo() BuildProcess {
	r.g.Ammo = 30
	return r
}

func (r *Rifle) SetDamage() BuildProcess {
	r.g.Damage = 10
	return r
}

func (r *Rifle) SetStructure() BuildProcess {
	r.g.Structure = "Rifle"
	return r
}

func (r *Rifle) GetGun() Gun {
	return r.g
}
