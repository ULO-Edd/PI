FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
WORKDIR /src

# Copy csproj and restore dependencies
COPY ["RegistrationsService.csproj", "./"]
RUN dotnet restore "RegistrationsService.csproj"

# Copy everything else and build
COPY . .
RUN dotnet publish "RegistrationsService.csproj" -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:10.0
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "RegistrationsService.dll"]
EXPOSE 80
