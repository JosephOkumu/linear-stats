package calcs

import "testing"

func TestLinearRegression(t *testing.T) {
	type args struct {
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		wantM float64
		wantC float64
	}{
		{
			name:  "Simple positive slope",
			args:  args{y: []float64{1, 2, 3, 4, 5}},
			wantM: 1,
			wantC: 1,
		},
		{
			name:  "Simple negative slope",
			args:  args{y: []float64{5, 4, 3, 2, 1}},
			wantM: -1,
			wantC: 5,
		},
		{
			name:  "Horizontal line",
			args:  args{y: []float64{3, 3, 3, 3, 3}},
			wantM: 0,
			wantC: 3,
		},
		{
			name:  "Vertical shift",
			args:  args{y: []float64{5, 5, 5, 5, 5}},
			wantM: 0,
			wantC: 5,
		},
		{
			name:  "Random points",
			args:  args{y: []float64{2, 4, 5, 4, 5}},
			wantM: 0.6,
			wantC: 2.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotC := LinearRegression(tt.args.y)
			if gotM != tt.wantM {
				t.Errorf("LinearRegression() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotC != tt.wantC {
				t.Errorf("LinearRegression() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestMean(t *testing.T) {
	type args struct {
		numslice []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Single element",
			args: args{numslice: []float64{15.0}},
			want: 15.0,
		},
		{
			name: "Multiple elements",
			args: args{numslice: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want: 3.0,
		},
		{
			name: "positive and negative elements",
			args: args{numslice: []float64{-1.0, -2.0, 3.0, 4.0}},
			want: 1.0,
		},
		{
			name: "All negative numbers",
			args: args{numslice: []float64{-1.0, -2.0, -3.0, -4.0, -5.0}},
			want: -3.0,
		},		
		{
			name: "Floating point numbers",
			args: args{numslice: []float64{1.5, 2.5, 3.5, 4.5}},
			want: 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.numslice); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPearsonsCorrelationCoefficient(t *testing.T) {
	tests := []struct {
		name     string
		y        []float64
		expected float64
	}{
		{
			name:     "Perfect Positive Correlation",
			y:        []float64{1, 2, 3, 4, 5},
			expected: 0.9999999999999998, // when rounded its 1
		},
		{
			name:     "Perfect Negative Correlation",
			y:        []float64{5, 4, 3, 2, 1},
			expected: -0.9999999999999998,
		},

		{
			name:     "Single Value",
			y:        []float64{1},
			expected: 0.0,
		},
		{
			name:     "All Same Values",
			y:        []float64{1, 1, 1, 1, 1},
			expected: 0.0, 
		},
		{
			name:     "Zero Variance in x",
			y:        []float64{1, 2, 3, 4, 5}, 
			expected: 0.9999999999999998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonCorrelation(tt.y); got != tt.expected {
				t.Errorf("PearsonsCorrelationCoefficient() = %v, want %v", got, tt.expected)
			}
		})
	}
}