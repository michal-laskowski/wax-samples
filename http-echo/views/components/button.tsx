export function Button(p: { title: string; fancy?: boolean }) {
    return p.fancy ? (
        <button>
            <marquee>{p.title}</marquee>
        </button>
    ) : (
        <button>{p.title}</button>
    )
}
