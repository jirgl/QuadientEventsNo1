import { Graph } from 'jirgl-data-structures';

type V = Graph.Vertex;

class EvaluatedVertex {
    vertex: V;
    prevEvaluatedVertex: EvaluatedVertex;
    evaluation: number;
    processed: boolean;

    constructor(vertex: V, evaluation: number) {
        this.vertex = vertex;
        this.evaluation = evaluation;
    }
}

export class Dijkstra {
    private vertices: EvaluatedVertex[];

    find(graph: Graph.Structure, startVertex: V, endVertex: V): V[] {
        this.init(graph, startVertex);
        return this.constructPath(this.evaluate(graph, startVertex, endVertex));
    }

    private init(graph: Graph.Structure, startVertex: V) {
        this.vertices = graph.getVertices().map((v) => {
            return {
                vertex: v,
                prevEvaluatedVertex: undefined,
                evaluation: v.id === startVertex.id ? 0 : undefined,
                processed: false
            };
        });
    }

    private getVertexWithBestEvaluation(vertices: EvaluatedVertex[]): EvaluatedVertex {
        let vertex: EvaluatedVertex = undefined;
        let bestEvaluation: number = Number.MAX_VALUE;

        vertices.forEach((v, i) => {
            if (v.evaluation < bestEvaluation && !v.processed) {
                bestEvaluation = v.evaluation;
                vertex = v;
            }
        });

        return vertex;
    }

    private evaluate(graph: Graph.Structure, startVertex: V, endVertex: V): EvaluatedVertex {
        let iteration = 0;
        let evaluatedVertex: EvaluatedVertex = undefined;
        for (let i = 0; i < this.vertices.length; i++) {
            evaluatedVertex = this.getVertexWithBestEvaluation(this.vertices);
            evaluatedVertex.processed = true;

            if (evaluatedVertex.vertex.id === endVertex.id) break;

            graph.getAdjacentEdges(evaluatedVertex.vertex).forEach((e) => {
                const nextVertex = e.v1.id === evaluatedVertex.vertex.id ? e.v2 : e.v1;
                const nextEvaluatedVertex = this.vertices.filter((v) => v.vertex.id === nextVertex.id)[0];
                const alt = evaluatedVertex.evaluation + e.evaluation;
                if (nextEvaluatedVertex.evaluation === undefined ||
                    alt < nextEvaluatedVertex.evaluation) {
                    nextEvaluatedVertex.evaluation = alt;
                    nextEvaluatedVertex.prevEvaluatedVertex = evaluatedVertex;
                }
            });
        }

        return evaluatedVertex;
    }

    private constructPath(endVertex: EvaluatedVertex): V[] {
        const path: V[] = [];
        let currentEvaluatedVertex = endVertex;
        while (currentEvaluatedVertex) {
            path.push(currentEvaluatedVertex.vertex);
            currentEvaluatedVertex = currentEvaluatedVertex.prevEvaluatedVertex;
        }

        return path.reverse();
    }
}