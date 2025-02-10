export function Layout(p: PropsWithChildren) {
    return (
        <>
            {wax.Raw(`<!DOCTYPE html>`)}
            <html>
                <head>
                    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
                    <script src="https://unpkg.com/htmx.org@1.9.3"></script>
                    <script src="https://unpkg.com/hyperscript.org@0.9.9"></script>
                    <link rel="stylesheet" href="https://unpkg.com/missing.css@1.1.3"></link>
                    {viewUtils.IsDev ? <script src="http://localhost:8181/live-reload.js"></script> : undefined}
                    <title>WAX - hello</title>
                </head>
                <body>
                    <h1>Hi there</h1>
                    <div>Layout content</div>
                    <div>{viewUtils.IsDev ? 'DEV RUN - live-reload enabled.' : 'Using embedded views'}</div>
                    <div>{viewUtils.IsDev && 'Edit and save HelloPage.tsx or any other TSX.'}</div>
                    <details open>
                        <summary>View content</summary>
                        <article>{p.children}</article>
                    </details>

                    <div>Other layout content</div>
                </body>
            </html>
        </>
    )
}
