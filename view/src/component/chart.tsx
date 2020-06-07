import { default as React, useEffect, useRef } from "react";
import { backgroundColor, borderColor } from "../assets/chart-common";
import { ResultDData } from "../store";
import Chart from 'chart.js'

interface ChartProps {
  type: "doughnut" | "bar" | 'line',
  data: Chart.ChartData
  options?: Chart.ChartOptions
  title?: string
}

function useChart ( param: ChartProps, dep: any[] ) {
  const canvasRef = useRef<HTMLCanvasElement>( null );
  const chartRef = useRef<Chart>();
  useEffect( () => {
    if ( canvasRef.current ) {
      const ctx = canvasRef.current.getContext( '2d' );
      if ( chartRef.current !== undefined && ctx !== null ) {
        chartRef.current.data = param.data;
        chartRef.current.update();
      } else {
        if ( ctx !== null ) {
          chartRef.current = new Chart( ctx, {
            type: param.type,
            data: param.data,
            options: {
              ...(param.options || {}),
              title: {
                display: param.type !== undefined,
                text: param.title || "",
              },
              scales: {
                yAxes: [ {
                  ticks: {
                    beginAtZero: true
                  }
                } ]
              },
              layout: {
                padding: {
                  left: 50,
                  right: 50,
                  top: 50,
                  bottom: 50
                }
              }
            }
          } )
        }
      }
    }
  }, dep );
  return canvasRef
}

interface BarChartProps {
  data: ResultDData,
  expense?: boolean
}

export function BarChart ( { data }: BarChartProps ) {
  const canvasRef = useChart( {
    type: "bar",
    data: {
      labels: data.children.map( v1 => v1.name ),
      datasets: [
        {
          label: "收入",
          data: data.children.map( v1 => v1.incomeMoney ),
          backgroundColor: backgroundColor[ 3 ],
          borderColor: borderColor[ 3 ],
          borderWidth: 1
        },
        {
          label: "支出",
          data: data.children.map( v1 => v1.expenseMoney ),
          backgroundColor: backgroundColor[ 5 ],
          borderColor: borderColor[ 5 ],
          borderWidth: 1
        }
      ]
    },
    title: "",
  }, [ data ] );
  return (
    <div className="chart-container">
      <canvas ref={ canvasRef }/>
    </div>
  )
}

export function BarChartTag ( { data, expense }: BarChartProps ) {
  const tags = expense ? data.tags.expense : data.tags.income;
  const canvasRef = useChart( {
    type: "bar",
    data: {
      labels: tags.map( v1 => v1.key ),
      datasets: [
        {
          label: expense ? "支出" : "收入",
          data: tags.map( v1 => v1.value ),
          backgroundColor: backgroundColor[ expense? 5 : 3 ],
          borderColor: borderColor[ expense? 5 : 3 ],
          borderWidth: 1
        }
      ]
    },
    title: "",
  }, [ data ] );
  return (
    <div className="chart-container">
      <canvas ref={ canvasRef }/>
    </div>
  )
}
