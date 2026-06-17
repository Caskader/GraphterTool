# Equation Graphter Frontend

A React-based frontend for the Graphter application that allows you to send mathematical equations to a Go backend and visualize the resulting points.

## Features

- 📝 Input equations in the custom format
- 🔄 Real-time communication with Go backend via POST requests
- 📊 Canvas-based visualization of equation points
- 📋 Table view showing all calculated points
- ⚠️ Error handling with user-friendly messages
- 🎨 Responsive design with smooth animations

## Getting Started

### Prerequisites
- Node.js and npm installed
- Go backend server running on `http://localhost:8080`

### Installation

1. Navigate to the frontend directory:
```bash
cd frontEnd
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

The application will open at `http://localhost:5173`

## Usage

1. **Enter an Equation**: Type your equation in the input field
   - Example: `((1x)^1 (+)^1 (1y)^1)^1`
   
2. **Submit**: Click "Graph Equation" button

3. **View Results**:
   - **Visualization**: Canvas chart showing all points plotted
   - **Points Table**: Scrollable list of all calculated points (X, Y coordinates)

## API Integration

The frontend communicates with the Go backend through:

**Endpoint**: `POST /api/equation`

**Request Body**:
```json
{
  "equation": "((1x)^1 (+)^1 (1y)^1)^1"
}
```

**Response**:
```json
{
  "status": "success",
  "points": [
    [0, 0],
    [1, 1],
    [2, 2],
    ...
  ]
}
```

## Building for Production

```bash
npm run build
```

The optimized build will be in the `dist/` directory.

## Troubleshooting

- **Connection Error**: Ensure the Go backend is running on port 8080
- **CORS Issues**: The backend should have CORS enabled for `*`
- **No Points Returned**: Check the equation format matches what the parser expects

## File Structure

- `src/App.jsx` - Main application component with all UI and logic
- `src/App.css` - Styling for the application
- `src/main.jsx` - React entry point
- `src/index.css` - Global styles
- `vite.config.js` - Vite configuration
