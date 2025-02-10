import { Button } from './components/button.tsx'
import { Layout } from './layout.tsx'

export default function Main(p: PageModel) {
    return (
        <Layout>
            <Button title="some button" />
            <Button title="fancy button" fancy />
            <div
                style={{
                    display: 'grid',
                    gridTemplateColumns: '1fr 1fr'
                }}
            >
                <span>Server started on: </span>
                <span>{p.ServerStartedOn}</span>
                <span>Response rendered on: </span>
                <span>{p.ServerRenderedOn}</span>
                <span>Date from view: </span>
                <span>{new Date().toISOString()}</span>
            </div>

            <div>
                <div>SomeString from model (your query string): {p.SomeString}</div>
                <div>Additional.IntValue: {p.Additional.IntValue}</div>
            </div>
        </Layout>
    )
}
