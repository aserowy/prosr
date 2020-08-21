// Generated with prosr1 by compiler 0.9, template csharp.tmpl 0.1

using Microsoft.AspNetCore.SignalR;
using Microsoft.AspNetCore.SignalR.Client;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Test.Package
{
	public interface ITestHubClient : ITestHubClientBase, IDisposable
	{
		Task InitializeHubConnection();
		Task CallRefreshFunctionTreeOnHub();
		Task CallTriggerOnHub();
		Task CallFunctionOnHub(FunctionRequest message);
	}

	public interface ITestHubClientBase
	{
		Task FunctionCalled(FunctionRequest message);
		Task FunctionTreeRefreshed();
		Task Function(FunctionResponse message);
	}

	public abstract class TestHubClientBase : ITestHubClient
	{
		private readonly IHubConnectionBuilder _hubConnectionBuilder;

		private bool _disposed = false;
		private HubConnection _connection;

		protected TestHubClientBase(IHubConnectionBuilder hubConnectionBuilder)
		{
			_hubConnectionBuilder = hubConnectionBuilder;
		}

		protected abstract Uri HubUrl { get; }

		public async Task InitializeHubConnection()
		{
			await GetConnection().ConfigureAwait(false);
		}

		public abstract Task FunctionCalled(FunctionRequest message);

		public abstract Task FunctionTreeRefreshed();

		public abstract Task Function(FunctionResponse message);

		public async Task CallRefreshFunctionTreeOnHub()
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("RefreshFunctionTree")
				.ConfigureAwait(false);
		}

		public async Task CallTriggerOnHub()
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("Trigger")
				.ConfigureAwait(false);
		}

		public async Task CallFunctionOnHub(FunctionRequest message)
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("Function", message)
				.ConfigureAwait(false);
		}

		private async Task<HubConnection> GetConnection()
		{
			if (_connection == null)
			{
				_connection = await CreateHubConnectionAsync().ConfigureAwait(false);
			}
		
			return _connection;
		}

		private async Task<HubConnection> CreateHubConnectionAsync()
		{
			var connection = _hubConnectionBuilder
				.WithUrl(HubUrl)
				.WithAutomaticReconnect(new[] { TimeSpan.Zero, TimeSpan.Zero, TimeSpan.FromSeconds(10) })
				.Build();

			BindClientMethods(ref connection);

			await connection
				.StartAsync()
				.ConfigureAwait(false);

			return connection;
		}

		private HubConnection BindClientMethods(ref HubConnection connection)
		{
			connection.On<FunctionRequest>("FunctionCalled", message => FunctionCalled(message));
			connection.On("FunctionTreeRefreshed", () => FunctionTreeRefreshed());
			connection.On<FunctionResponse>("Function", message => Function(message));

			return connection;
		}

		public void Dispose()
		{
			Dispose(true);
			GC.SuppressFinalize(this);
		}
		
		protected virtual void Dispose(bool disposing)
		{
			if (!_disposed)
			{
				if (disposing)
				{
					if (!(_connection is null))
					{
						var result = _connection
							.DisposeAsync()
							.Wait(1000);

						if (!result)
						{
							throw new TimeoutException($"Timeout of one second reached while closing hub connection for {HubUrl}!");
						}
					}
				}

				_disposed = true;
			}
		}
		
		~TestHubClientBase()
		{
			Dispose(false);
		}
	}

	public abstract class TestHubBase : Hub<ITestHubClientBase>
	{
		protected Task SendFunctionCalledOnAllAsync(FunctionRequest message)
		{
			return Clients.All.FunctionCalled(message);
		}

		public async Task RefreshFunctionTree()
		{
			await HandleRefreshFunctionTreeAsync().ConfigureAwait(false);

			await Clients
				.Caller
				.FunctionTreeRefreshed()
				.ConfigureAwait(false);
		}

		protected abstract Task HandleRefreshFunctionTreeAsync();
		
		public async Task Trigger()
		{
			await HandleTriggerAsync().ConfigureAwait(false);
		}

		protected abstract Task HandleTriggerAsync();
		
		public async Task Function(FunctionRequest message)
		{
			var result = await HandleFunctionAsync(message).ConfigureAwait(false);

			await Clients
				.Caller
				.Function(result)
				.ConfigureAwait(false);
		}

		protected abstract Task<FunctionResponse> HandleFunctionAsync(FunctionRequest message);
		
	}

	public class FunctionRequest
	{
		public int result_per_page { get; set; }
		public int page_number { get; set; }
		public string name { get; set; }
	}

	public class FunctionResponse
	{
		public IEnumerable<Test.Package2.Result> result { get; set; }
	}
}
	
namespace Test.Package2
{
	public class Result
	{
		public string snippets { get; set; }
		public string title { get; set; }
		public string url { get; set; }
	}
}
	