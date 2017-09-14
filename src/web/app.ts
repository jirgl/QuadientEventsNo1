import * as b from 'bobril';
import { Button } from 'bobril-m';
import { store } from './store';

export interface IAppData {
}

interface IAppCtx extends b.IBobrilCtx {
    data: IAppData;
}

export const App = b.createComponent<IAppData>({
    render(ctx: IAppCtx, me: b.IBobrilNode) {
        const d = ctx.data;
        me.children = Button({ action: store.acquireTask }, 'acquire task');
    }
});
