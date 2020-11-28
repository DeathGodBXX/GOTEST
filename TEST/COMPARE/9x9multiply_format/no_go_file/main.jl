using BenchmarkTools

@time function test()
    for i = 1:9
        for j = 1:i
            print(i,"x",j,"=",i*j,"\t")
        end
        println()
    end
end

for i = 1:3
    test()
    # println(i)  
end