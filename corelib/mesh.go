package corelib

type MeshInfo struct {
	Id            int
	Name          string
	Attributes    []int
	Indexes       []int
	NormalIndexes []int
	TexIndexes    []int
	Owner         *MeshData
}

type MeshData struct {
	Vertices  []Vector3
	Normals   []Vector3
	TexCoords []Vector2
	Meshes    []MeshInfo
}
