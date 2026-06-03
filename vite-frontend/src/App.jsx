import { useState } from 'react'
import './App.css'

function App() {
  // States to hold the input name, backend response, and any errors
  const [name, setName] = useState('')
  const [backendData, setBackendData] = useState(null)
  const [error, setError] = useState(null)

  // Function to handle sending the GET request
  const handleGetData = async () => {
    try {
      setError(null) // Reset errors
      
      // Build URL with query parameter (defaults to 'React User' if input is empty)
      const targetName = name.trim() || 'React User'
      const response = await fetch(`http://localhost:8080/api/data?name=${targetName}`)
      
      if (!response.ok) {
        throw new Error('Failed to fetch data from Go backend')
      }

      const data = await response.json()
      setBackendData(data) // Save the response inside state
    } catch (err) {
      setError(err.message)
      setBackendData(null)
    }
  }

  return (
    <div style={styles.container}>
      <h2>Go + React GET Data Exchange</h2>
      
      <div style={styles.card}>
        <input 
          type="text" 
          placeholder="Enter a name to send..." 
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={styles.input}
        />
        <button onClick={handleGetData} style={styles.button}>
          Send GET Request
        </button>
      </div>

      {error && <div style={styles.error}>Error: {error}</div>}

      {backendData && (
        <div style={styles.responseBox}>
          <h3>Backend Response Received:</h3>
          <p><strong>Message:</strong> {backendData.message}</p>
          <p><strong>Status:</strong> <span style={styles.statusBadge}>{backendData.status}</span></p>
        </div>
      )}
    </div>
  )
}

// Simple inline styling
const styles = {
  container: { fontFamily: 'Arial, sans-serif', padding: '40px', maxWidth: '500px', margin: '0 auto', textAlign: 'center' },
  card: { background: '#f9f9f9', padding: '20px', borderRadius: '8px', boxShadow: '0 4px 6px rgba(0,0,0,0.1)', marginBottom: '20px' },
  input: { padding: '10px', fontSize: '16px', borderRadius: '4px', border: '1px solid #ccc', marginRight: '10px', width: '60%' },
  button: { padding: '10px 15px', fontSize: '16px', background: '#61dafb', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold' },
  responseBox: { background: '#e6f7ff', borderLeft: '5px solid #1890ff', padding: '15px', borderRadius: '4px', textAlign: 'left' },
  statusBadge: { background: '#52c41a', color: 'white', padding: '2px 8px', borderRadius: '4px', fontSize: '14px' },
  error: { color: 'red', background: '#fff2f0', padding: '10px', borderRadius: '4px', border: '1px solid #ffccc7' }
}

export default App