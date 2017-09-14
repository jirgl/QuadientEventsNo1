import * as api from './apiBase';

export interface IPosition {
    x: number;
    y: number;
}

export interface ITask {
    id: number;
    startedTimestamp: number;
    map: {
        areas: string[]
    };
    astroants: IPosition;
    sugar: IPosition;
}

export interface ITaskResult {
    valid: boolean;
    inTime: boolean;
    message: string;
}

export function getTask(): Promise<ITask> {
    return api.callGet('task');
}

export function submitTask(id: number, answer: string): Promise<ITaskResult> {
    return api.callPut('task/' + id, { path: answer });
}
