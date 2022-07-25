package graph

const (
	Vertex_amonut = 15
)

var Cascade_Position = [][]int{
	{10, 5, 0},
	{11, 6, 1},
	{12, 7, 2},
	{13, 8, 3},
	{14, 9, 4},
}

func init() {
	//NewGraph().Init_Vertex_Graph()
}

func each_number_connect() [Vertex_amonut][Vertex_amonut]int {
	var vertex_connet_table [Vertex_amonut][Vertex_amonut]int
	connet := map[int][]int{
		0:  {1, 5},
		1:  {0, 2, 6},
		2:  {1, 3, 7},
		3:  {2, 4, 8},
		4:  {3, 9},
		5:  {0, 6, 10},
		6:  {1, 5, 7, 11},
		7:  {2, 6, 8, 12},
		8:  {3, 7, 9, 13},
		9:  {4, 8, 14},
		10: {4, 11},
		11: {6, 10, 12},
		12: {7, 11, 13},
		13: {8, 12, 14},
		14: {9, 13},
	}

	for v := 0; v < Vertex_amonut; v++ {
		for _, cor_v := range connet[v] {
			vertex_connet_table[v][cor_v] = 1
		}
	}

	//fmt.Println(vertex_connet_table)
	return vertex_connet_table
}

func (g *Graph) Init_Vertex() {
	for v := 0; v < Vertex_amonut; v++ {
		g.NewVertex(v)
	}
}

func (g *Graph) Init_Vertex_Graph() {
	table := each_number_connect()
	for v := 0; v < Vertex_amonut; v++ {
		for k := v; k < Vertex_amonut; k++ {
			if table[v][k] == 1 {
				g.AddEdge(v, k)
			}
		}
	}

}
