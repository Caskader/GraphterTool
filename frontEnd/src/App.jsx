import { useState, useRef, useEffect } from 'react'
import './App.css'

function App() {
  const [equation, setEquation] = useState('')
  const [points, setPoints] = useState([])
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(false)
  const canvasRef = useRef(null)

  // Function to handle sending the equation
  const handleSubmit = async (e) => {
    e.preventDefault()
    setError(null)


    if (!equation.trim()) {
      setError('Please enter an equation')
      return
    }

    setLoading(true)
    try {
      const response = await fetch('http://localhost:8080/api/equation', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ equation: equation.trim() }),
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()

      if (data.status === 'error') {
        throw new Error(data.error || 'Failed to process equation')
      }

      setPoints(data.points || [])
    } catch (err) {
      setError(err.message)
      setPoints([])
    } finally {
      setLoading(false)
    }
  }

  // Draw canvas plot when points change
  useEffect(() => {
    if (points.length === 0 || !canvasRef.current) return

    const canvas = canvasRef.current
    const ctx = canvas.getContext('2d')
    const width = canvas.width
    const height = canvas.height
    const padding = 40

    // Find min/max values for scaling
    let minX = Infinity, maxX = -Infinity
    let minY = Infinity, maxY = -Infinity

    points.forEach(([x, y]) => {
      if (x < minX) minX = x
      if (x > maxX) maxX = x
      if (y < minY) minY = y
      if (y > maxY) maxY = y
    })

    // Add padding to min/max
    const rangeX = maxX - minX || 1
    const rangeY = maxY - minY || 1
    minX -= rangeX * 0.1
    maxX += rangeX * 0.1
    minY -= rangeY * 0.1
    maxY += rangeY * 0.1

    const scaleX = (width - 2 * padding) / (maxX - minX)
    const scaleY = (height - 2 * padding) / (maxY - minY)

    // Helper to convert data coordinates to canvas coordinates
    const toCanvasX = (x) => padding + (x - minX) * scaleX
    const toCanvasY = (y) => height - padding - (y - minY) * scaleY

    // Clear canvas
    ctx.fillStyle = 'white'
    ctx.fillRect(0, 0, width, height)

    // Draw grid
    ctx.strokeStyle = '#e0e0e0'
    ctx.lineWidth = 0.5
    for (let x = Math.ceil(minX); x <= Math.floor(maxX); x++) {
      if (x % Math.ceil(rangeX / 10) === 0) {
        ctx.beginPath()
        ctx.moveTo(toCanvasX(x), padding)
        ctx.lineTo(toCanvasX(x), height - padding)
        ctx.stroke()
      }
    }
    for (let y = Math.ceil(minY); y <= Math.floor(maxY); y++) {
      if (y % Math.ceil(rangeY / 10) === 0) {
        ctx.beginPath()
        ctx.moveTo(padding, toCanvasY(y))
        ctx.lineTo(width - padding, toCanvasY(y))
        ctx.stroke()
      }
    }

    // Draw axes
    ctx.strokeStyle = 'black'
    ctx.lineWidth = 2
    ctx.beginPath()
    ctx.moveTo(padding, height - padding)
    ctx.lineTo(width - padding, height - padding)
    ctx.stroke()

    ctx.beginPath()
    ctx.moveTo(padding, padding)
    ctx.lineTo(padding, height - padding)
    ctx.stroke()

    // Draw axis labels
    ctx.fillStyle = 'black'
    ctx.font = '12px Arial'
    ctx.textAlign = 'center'
    ctx.fillText('X', width - 20, height - 10)
    ctx.save()
    ctx.translate(10, height / 2)
    ctx.rotate(-Math.PI / 2)
    ctx.fillText('Y', 0, 0)
    ctx.restore()

    // Draw points
    ctx.fillStyle = '#ff6b6b'
    points.forEach(([x, y]) => {
      const canvasX = toCanvasX(x)
      const canvasY = toCanvasY(y)
      ctx.beginPath()
      ctx.arc(canvasX, canvasY, 3, 0, 2 * Math.PI)
      ctx.fill()
    })
  }, [points])

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h1>Equation Graphter</h1>
        <p>Send equations to the Go server and visualize the points</p>
      </div>

      <div style={styles.inputSection}>
        <form onSubmit={handleSubmit} style={styles.form}>
          <input
            type="text"
            placeholder="Enter equation (e.g., ((1x)^1 (+)^1 (1y)^1)^1)"
            value={equation}
            onChange={(e) => setEquation(e.target.value)}
            style={styles.input}
            disabled={loading}
          />
          <button type="submit" style={styles.button} disabled={loading}>
            {loading ? 'Processing...' : 'Graph Equation'}
          </button>
        </form>
      </div>

      {error && <div style={styles.error}>⚠️ Error: {error}</div>}

      <div style={styles.content}>
        {points.length > 0 && (
          <>
            <div style={styles.canvasSection}>
              <h2>Visualization</h2>
              <canvas
                ref={canvasRef}
                width={600}
                height={600}
                style={styles.canvas}
              />
            </div>

            <div style={styles.pointsSection}>
              <h2>Points ({points.length} total)</h2>
              <div style={styles.pointsTableContainer}>
                <table style={styles.table}>
                  <thead>
                    <tr style={styles.tableHeader}>
                      <th style={styles.th}>X</th>
                      <th style={styles.th}>Y</th>
                    </tr>
                  </thead>
                  <tbody>
                    {points.slice(0, 20).map((point, idx) => (
                      <tr key={idx} style={idx % 2 === 0 ? styles.evenRow : styles.oddRow}>
                        <td style={styles.td}>{point[0]}</td>
                        <td style={styles.td}>{point[1]}</td>
                      </tr>
                    ))}
                    {points.length > 20 && (
                      <tr>
                        <td colSpan="2" style={styles.moreRow}>
                          ... and {points.length - 20} more points
                        </td>
                      </tr>
                    )}
                  </tbody>
                </table>
              </div>
            </div>
          </>
        )}
        {!loading && points.length === 0 && !error && (
          <div style={styles.placeholder}>
            <p>Enter an equation and click "Graph Equation" to see results</p>
          </div>
        )}
      </div>
    </div>
  )
}

const styles = {
  container: {
    maxWidth: '1200px',
    margin: '0 auto',
    padding: '20px',
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif',
    backgroundColor: '#f5f5f5',
    minHeight: '100vh',
  },
  header: {
    textAlign: 'center',
    marginBottom: '30px',
    padding: '20px',
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
  },
  inputSection: {
    marginBottom: '30px',
  },
  form: {
    display: 'flex',
    gap: '10px',
    padding: '20px',
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
  },
  input: {
    flex: 1,
    padding: '12px 16px',
    fontSize: '14px',
    border: '2px solid #ddd',
    borderRadius: '4px',
    fontFamily: 'monospace',
    transition: 'border-color 0.3s',
  },
  button: {
    padding: '12px 24px',
    fontSize: '14px',
    fontWeight: '600',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '4px',
    cursor: 'pointer',
    transition: 'background-color 0.3s',
  },
  error: {
    padding: '15px',
    marginBottom: '20px',
    backgroundColor: '#ffe6e6',
    color: '#d32f2f',
    borderRadius: '4px',
    border: '1px solid #ff9999',
  },
  content: {
    display: 'flex',
    gap: '30px',
    flexWrap: 'wrap',
  },
  canvasSection: {
    flex: 1,
    minWidth: '300px',
    backgroundColor: 'white',
    borderRadius: '8px',
    padding: '20px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
  },
  canvas: {
    border: '1px solid #ddd',
    borderRadius: '4px',
    maxWidth: '100%',
    height: 'auto',
    display: 'block',
    marginTop: '10px',
  },
  pointsSection: {
    flex: 1,
    minWidth: '300px',
    backgroundColor: 'white',
    borderRadius: '8px',
    padding: '20px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
  },
  pointsTableContainer: {
    maxHeight: '600px',
    overflowY: 'auto',
    border: '1px solid #ddd',
    borderRadius: '4px',
  },
  table: {
    width: '100%',
    borderCollapse: 'collapse',
    fontSize: '13px',
  },
  tableHeader: {
    backgroundColor: '#f0f0f0',
    position: 'sticky',
    top: 0,
  },
  th: {
    padding: '10px',
    textAlign: 'center',
    fontWeight: '600',
    borderBottom: '2px solid #ddd',
  },
  td: {
    padding: '10px',
    textAlign: 'center',
    borderBottom: '1px solid #eee',
  },
  evenRow: {
    backgroundColor: '#fafafa',
  },
  oddRow: {
    backgroundColor: 'white',
  },
  moreRow: {
    textAlign: 'center',
    padding: '10px',
    color: '#666',
    fontStyle: 'italic',
  },
  placeholder: {
    gridColumn: '1 / -1',
    textAlign: 'center',
    padding: '60px 20px',
    color: '#999',
    fontSize: '16px',
  },
}

export default App