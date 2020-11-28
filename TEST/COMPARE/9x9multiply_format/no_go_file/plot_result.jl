using Plots

y=[
    0.04	0.085	0.103	0.123	0.163;
    0.043	0.072	0.065	0.065	0.129;
    0.038	0.063	0.076	0.094	0.116;
    0.04	0.077	0.054	0.037	0.072;
    0.042	0.086	0.091	0.05	0.056;
] # 竖列画图 "julia" "go" "python" "c" "c++"

scatter(y,label=false)
plot!(y,label=["julia" "go" "python" "c" "c++"],lw=3)
savefig("plot_result.pdf")
savefig("plot_result.png")

