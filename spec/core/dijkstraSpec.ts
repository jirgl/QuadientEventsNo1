import { Graph } from 'jirgl-data-structures';
import { Dijkstra } from '../../src/core/dijkstra';

describe('Dijkstra algorithm', () => {

    const v1: Graph.Vertex = { id: 'v1', position: [0, 0] };
    const v2: Graph.Vertex = { id: 'v2', position: [10, 0] };
    const v3: Graph.Vertex = { id: 'v3', position: [15, 5] };
    const v4: Graph.Vertex = { id: 'v4', position: [10, 10] };
    const v5: Graph.Vertex = { id: 'v5', position: [0, 10] };
    const v6: Graph.Vertex = { id: 'v6', position: [0, 20] };
    const e1 = { evaluation: 10, v1: v1, v2: v2 };
    const e2 = { evaluation: 10, v1: v2, v2: v3 };
    const e3 = { evaluation: 9, v1: v3, v2: v4 };
    const e4 = { evaluation: 10, v1: v4, v2: v1 };
    const e5 = { evaluation: 5, v1: v4, v2: v5 };
    const e6 = { evaluation: 10, v1: v5, v2: v1 };
    const e7 = { evaluation: 4, v1: v5, v2: v6 };
    const e8 = { evaluation: 10, v1: v6, v2: v4 };

    it('find path', () => {
        const g = new Graph.Structure();
        g.addEdge(e1);
        g.addEdge(e2);
        g.addEdge(e3);
        g.addEdge(e4);
        g.addEdge(e5);
        g.addEdge(e6);
        g.addEdge(e7);
        g.addEdge(e8);
        const d = new Dijkstra();
        expect(d.find(g, v1, v3)).toEqual([v1, v4, v3]);
        expect(d.find(g, v2, v4)).toEqual([v2, v3, v4]);
        expect(d.find(g, v3, v6)).toEqual([v3, v4, v5, v6]);
    });

});
